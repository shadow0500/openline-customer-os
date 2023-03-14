import {
  GetContactsSuggestionsQuery,
  useGetContactsSuggestionsQuery,
} from './types';
import { ApolloError } from 'apollo-client';

interface Props {
  value: string;
}

interface Result {
  data: GetContactsSuggestionsQuery['contacts'] | null | undefined;
  loading: boolean;
  error: ApolloError | null;
}
export const useContactSuggestions = ({ value }: Props): Result => {
  const { data, loading, error } = useGetContactsSuggestionsQuery({
    variables: { value },
  });

  if (loading) {
    return {
      loading: true,
      error: null,
      data: null,
    };
  }

  if (error) {
    return {
      error,
      loading: false,
      data: null,
    };
  }

  return {
    data: data?.contacts ?? null,
    loading,
    error: null,
  };
};
