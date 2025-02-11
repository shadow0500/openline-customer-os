import { ApolloError } from '@apollo/client';
import { GetTagsQuery, useGetTagsQuery } from './types';
import { toast } from 'react-toastify';

interface Result {
  tags: GetTagsQuery['tags'] | undefined | null;
  loading: boolean;
  error: ApolloError | null;
}
export const useTags = (): Result => {
  const { data, loading, error } = useGetTagsQuery();

  if (loading) {
    return {
      loading: true,
      error: null,
      tags: null,
    };
  }

  if (error) {
    toast.error('Something went wrong while loading tags', {
      toastId: `tags-loading-error`,
    });
    return {
      error,
      loading: false,
      tags: null,
    };
  }

  return {
    tags: data?.tags,
    loading,
    error: null,
  };
};
