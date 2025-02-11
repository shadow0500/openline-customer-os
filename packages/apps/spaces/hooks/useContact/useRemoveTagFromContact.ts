import { ContactTagInput } from './types';
import {
  RemoveTagFromContactMutation,
  useRemoveTagFromContactMutation,
} from '../../graphQL/__generated__/generated';
import { toast } from 'react-toastify';

interface Result {
  onRemoveTagFromContact: (
    input: Omit<ContactTagInput, 'contactId'>,
  ) => Promise<RemoveTagFromContactMutation['contact_RemoveTagById'] | null>;
}
export const useRemoveTagFromContact = ({
  contactId,
}: {
  contactId: string;
}): Result => {
  const [removeTagFromContactMutation, { loading, error, data }] =
    useRemoveTagFromContactMutation();

  const handleRemoveTagFromContact: Result['onRemoveTagFromContact'] = async (
    contactTagInput,
  ) => {
    try {
      const response = await removeTagFromContactMutation({
        variables: { input: { ...contactTagInput, contactId } },
        update(cache) {
          const normalizedId = cache.identify({
            id: contactTagInput.tagId,
            __typename: 'Tag',
          });
          cache.evict({ id: normalizedId });
          cache.gc();
        },
      });
      return response.data?.contact_RemoveTagById ?? null;
    } catch (err) {
      console.error(err);
      toast.error(
        'Something went wrong while deleting tag. Please contact us or try again later',
        {
          toastId: `contact-tag-${contactId}-${contactTagInput.tagId}-delete-error`,
        },
      );
      return null;
    }
  };

  return {
    onRemoveTagFromContact: handleRemoveTagFromContact,
  };
};
