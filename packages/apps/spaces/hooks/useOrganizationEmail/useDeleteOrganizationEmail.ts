import {
  useRemoveEmailFromOrganizationMutation,
  RemoveEmailFromOrganizationMutation,
} from './types';
import { toast } from 'react-toastify';

interface Props {
  organizationId: string;
}

interface Result {
  onRemoveEmailFromOrganization: (
    emailId: string,
  ) => Promise<
    | RemoveEmailFromOrganizationMutation['emailRemoveFromOrganizationById']
    | null
  >;
}
export const useRemoveEmailFromOrganizationEmail = ({
  organizationId,
}: Props): Result => {
  const [removeEmailFromOrganizationMutation, { loading, error, data }] =
    useRemoveEmailFromOrganizationMutation();

  const handleRemoveEmailFromOrganization: Result['onRemoveEmailFromOrganization'] =
    async (emailId) => {
      try {
        const response = await removeEmailFromOrganizationMutation({
          variables: { organizationId, id: emailId },
          refetchQueries: ['GetOrganizationCommunicationChannels'],

          update(cache) {
            const normalizedId = cache.identify({
              id: emailId,
              __typename: 'Email',
            });
            cache.evict({ id: normalizedId });
            cache.gc();
          },
        });
        return response.data?.emailRemoveFromOrganizationById ?? null;
      } catch (err) {
        toast.error(
          'Something went wrong while deleting email! Please contact us or try again later',
          {
            toastId: `email-${emailId}-delete-error`,
          },
        );
        console.error(err);
        return null;
      }
    };

  return {
    onRemoveEmailFromOrganization: handleRemoveEmailFromOrganization,
  };
};
