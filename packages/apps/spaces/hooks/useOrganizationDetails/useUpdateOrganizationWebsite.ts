import {
  GetOrganizationDetailsQuery,
  UpdateOrganizationWebsiteMutation,
  useUpdateOrganizationWebsiteMutation,
} from './types';
import {
  GetContactPersonalDetailsWithOrganizationsDocument,
  GetOrganizationDetailsDocument,
  OrganizationUpdateInput,
} from '../../graphQL/__generated__/generated';
import { ApolloCache } from '@apollo/client/cache';
import client from '../../apollo-client';
import { toast } from 'react-toastify';

interface Props {
  organizationId: string;
}

interface Result {
  onUpdateOrganizationWebsite: (
    input: Omit<OrganizationUpdateInput, 'id'>,
  ) => Promise<UpdateOrganizationWebsiteMutation['organization_Update'] | null>;
}
export const useUpdateOrganizationWebsite = ({
  organizationId,
}: Props): Result => {
  const [updateOrganizationMutation, { loading, error, data }] =
    useUpdateOrganizationWebsiteMutation();

  const handleUpdateCacheAfterUpdatingOrganization = (
    cache: ApolloCache<any>,
    { data: { organization_Update } }: any,
  ) => {
    const data: GetOrganizationDetailsQuery | null = client.readQuery({
      query: GetOrganizationDetailsDocument,
      variables: {
        id: organizationId,
      },
    });

    if (data === null) {
      client.writeQuery({
        query: GetOrganizationDetailsDocument,
        data: {
          organization: {
            id: organizationId,
            ...organization_Update,
          },
          variables: { id: organizationId },
        },
      });
    }

    client.writeQuery({
      query: GetOrganizationDetailsDocument,
      data: {
        organization: {
          id: organizationId,
          ...data?.organization,
          website: organization_Update.website,
        },
      },
      variables: {
        id: organizationId,
      },
    });
  };

  const handleUpdateOrganizationWebsite: Result['onUpdateOrganizationWebsite'] =
    async (input) => {
      try {
        const response = await updateOrganizationMutation({
          variables: { input: { ...input, id: organizationId } },
          update: handleUpdateCacheAfterUpdatingOrganization,
        });
        return response.data?.organization_Update ?? null;
      } catch (err) {
        toast.error(
          'Something went wrong while updating organization website. Please contact us or try again later',
          {
            toastId: `org-website-${organizationId}-update-error`,
          },
        );
        console.error(err);
        return null;
      }
    };

  return {
    onUpdateOrganizationWebsite: handleUpdateOrganizationWebsite,
  };
};
