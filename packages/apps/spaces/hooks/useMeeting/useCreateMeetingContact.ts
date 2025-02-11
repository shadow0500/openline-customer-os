import {
  GetContactTimelineDocument,
  NOW_DATE,
  Result,
  useCreateMeetingMutation,
} from './types';
import { toast } from 'react-toastify';
import { ApolloCache } from '@apollo/client/cache';
import {
  DataSource,
  GetContactTimelineQuery,
} from '../../graphQL/__generated__/generated';
import client from '../../apollo-client';

export interface Props {
  contactId?: string;
}
export const useCreateMeetingFromContact = ({ contactId }: Props): Result => {
  const [createMeetingMutation, { loading, error, data }] =
    useCreateMeetingMutation();

  const handleUpdateCacheAfterAddingMeeting = (
    cache: ApolloCache<any>,
    { data: { meeting_Create } }: any,
  ) => {
    const data: GetContactTimelineQuery | null = client.readQuery({
      query: GetContactTimelineDocument,
      variables: {
        contactId,
        from: NOW_DATE,
        size: 10,
      },
    });

    const newMeeting = {
      createdAt: new Date(),
      meetingStartedAt: new Date(),
      meetingEndedAt: new Date(),
      agendaContentType: 'text/html',
      meetingCreatedBy: meeting_Create.createdBy,
      describedBy: [],
      includes: [],
      events: [],
      recording: null,
      id: meeting_Create.id,
      source: DataSource.Openline,
      ...meeting_Create,
    };

    if (data === null) {
      client.writeQuery({
        query: GetContactTimelineDocument,
        data: {
          contact: {
            contactId,
            timelineEvents: [newMeeting],
          },
          variables: { contactId, from: NOW_DATE, size: 10 },
        },
      });
      return;
    }

    const newData = {
      contact: {
        ...data.contact,
        timelineEvents: [newMeeting],
      },
    };

    client.writeQuery({
      query: GetContactTimelineDocument,
      data: newData,
      variables: {
        contactId,
        from: NOW_DATE,
        size: 10,
      },
    });
  };

  const handleCreateMeetingFromContact: Result['onCreateMeeting'] = async (
    userId,
  ) => {
    try {
      const response = await createMeetingMutation({
        variables: {
          meeting: {
            createdBy: [{ userId: userId }],
            attendedBy: [{ contactId: contactId }],
            appSource: 'OPENLINE',
            name: '',
            startedAt: new Date().toISOString(),
            endedAt: new Date().toISOString(),
            agenda: `<p>INTRODUCTION</p>
                     <p>DISCUSSION</p>
                     <p>NEXT STEPS</p>
                     `,
            agendaContentType: 'text/html',
            note: { html: '<p>Notes:</p>', appSource: 'OPENLINE' },
          },
        },

        update: handleUpdateCacheAfterAddingMeeting,
      });

      if (response.data?.meeting_Create.id) {
        toast.success(`Added draft meeting to the timeline`, {
          toastId: `draft-meeting-added-${response.data?.meeting_Create.id}`,
        });
      }

      return response.data?.meeting_Create ?? null;
    } catch (err) {
      console.error(err);
      toast.error(
        `Something went wrong while adding draft meeting to the timeline`,
      );
      return null;
    }
  };

  return {
    onCreateMeeting: handleCreateMeetingFromContact,
  };
};
