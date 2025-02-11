import { ApolloError, NetworkStatus } from '@apollo/client';
import { GetContactTimelineQuery, useGetContactTimelineQuery } from './types';
import { getContactDisplayName } from '../../utils';

interface Props {
  contactId: string;
}

interface Result {
  //@ts-expect-error fixme
  data: GetContactTimelineQuery['contact']['timelineEvents'] | null | undefined;
  loading: boolean;
  error: ApolloError | null;
  fetchMore: (data: { variables: any }) => void;
  variables: any;
  networkStatus?: NetworkStatus;
  contactName: string;
}

const NOW_DATE = new Date().toISOString();

export const useContactTimeline = ({ contactId }: Props): Result => {
  const { data, loading, error, variables, fetchMore, networkStatus } =
    useGetContactTimelineQuery({
      variables: {
        contactId,
        from: NOW_DATE,
        size: 15,
      },
      fetchPolicy: 'cache-and-network',
      nextFetchPolicy: 'cache-first',
      notifyOnNetworkStatusChange: true,
    });
  const timelineEvents = [...(data?.contact?.timelineEvents || [])].sort(
    (a, b) => {
      return (
        //@ts-expect-error fixme
        Date.parse(a?.createdAt || a?.startedAt) -
        //@ts-expect-error fixme
        Date.parse(b?.createdAt || b?.startedAt)
      );
    },
  );

  if (loading) {
    return {
      loading: true,
      error: null,
      // @ts-expect-error fixme
      contactName: data?.contact ? getContactDisplayName(data?.contact) : '',
      data: timelineEvents,
      fetchMore,
      variables: variables,
      networkStatus,
    };
  }

  if (error) {
    return {
      error,
      loading: false,
      variables: variables,
      networkStatus,
      data: null,
      contactName: '',
      fetchMore,
    };
  }

  return {
    data: timelineEvents,
    // @ts-expect-error fixme
    contactName: data?.contact ? getContactDisplayName(data?.contact) : '',
    fetchMore,
    loading,
    error: null,
    variables: variables,
    networkStatus,
  };
};
