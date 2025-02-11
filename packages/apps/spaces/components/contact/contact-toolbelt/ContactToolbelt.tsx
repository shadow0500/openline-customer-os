import React from 'react';

import { useCreateContactNote } from '@spaces/hooks/useNote';
import { useCreateMeetingFromContact } from '@spaces/hooks/useMeeting';
import { TimelineToolbelt } from '@spaces/molecules/timeline-toolbelt';
import { useRecoilState, useRecoilValue } from 'recoil';
import { contactNewItemsToEdit, userData } from '../../../state';
import { useUser } from '@spaces/hooks/useUser';
import { toast } from 'react-toastify';

interface ToolbeltProps {
  contactId: string;
  isSkewed: boolean;
}

export const ContactToolbelt: React.FC<ToolbeltProps> = ({
  contactId,
  isSkewed,
}) => {
  const [itemsInEditMode, setItemToEditMode] = useRecoilState(
    contactNewItemsToEdit,
  );
  const { identity: userEmail } = useRecoilValue(userData);
  const { data, loading, error } = useUser({ email: userEmail });
  const { onCreateContactNote, saving } = useCreateContactNote({ contactId });
  const { onCreateMeeting } = useCreateMeetingFromContact({ contactId });

  const handleCreateNote = (data: any) =>
    onCreateContactNote(data).then((response) => {
      if (response?.id) {
        setItemToEditMode({
          timelineEvents: [
            ...itemsInEditMode.timelineEvents,
            { id: response.id },
          ],
        });
      }
    });

  const handleCreateMeeting = () => {
    if (!data?.id) {
      toast.error('Meeting could not be created, please try again later');
      return;
    }
    return onCreateMeeting(data?.id);
  };

  return (
    <TimelineToolbelt
      showPhoneCallButton
      onCreateMeeting={handleCreateMeeting}
      onCreateNote={handleCreateNote}
      isSkewed={isSkewed}
    />
  );
};
