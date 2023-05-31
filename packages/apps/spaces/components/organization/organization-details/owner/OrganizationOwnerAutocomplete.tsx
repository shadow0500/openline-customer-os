import React, { useEffect, useState } from 'react';
import { useOrganizationOwner } from '@spaces/hooks/useOrganizationOwner';
import { useLinkOrganizationOwner } from '@spaces/hooks/useOrganizationOwner/useLinkOrganizationOwner';
import { useUserSuggestionsList } from '@spaces/hooks/useUser';
import { useUnlinkOrganizationOwner } from '@spaces/hooks/useOrganizationOwner/useUnlinkOrganizationOwner';
import { Autocomplete } from '@spaces/atoms/new-autocomplete';
import { SuggestionItem } from '@spaces/atoms/new-autocomplete/Autocomplete';
import { useRecoilValue } from 'recoil';
import { ownerListData } from '../../../../state/userData';
import Fuse from 'fuse.js'

interface OrganizationOwnerProps {
  id: string;
  editMode: boolean;
  switchEditMode?: () => void;
}

export const OrganizationOwnerAutocomplete: React.FC<
  OrganizationOwnerProps
> = ({ id, editMode, switchEditMode }) => {
  const [ownerListMatch, setOwnerListMatch] = useState<Array<SuggestionItem>>(
    [],
  );
  const [ownerListFuzzy, setOwnerListFuzzy] = useState<Array<SuggestionItem>>(
    [],
  );
  const [userId, setUserId] = React.useState<string>('');
  const [inputValue, setInputValue] = React.useState<string>('');

  const [loadingOwnerSuggestions, setLoadingOwnerSuggestions] =
    useState<boolean>(false);
  const {ownerList} = useRecoilValue(ownerListData);
  const fuse = new Fuse(ownerList, {keys: ['firstName', 'lastName']});

  const mapOwnerToSuggestionItem = (owner: any) => {
    return ({ label: `${owner.firstName} ${owner.lastName}`, value: owner.id });
  }
  const searchOwners = (filter: string) => {
    setLoadingOwnerSuggestions(true);
    const ownersContains = ownerList.filter((owner: any) => owner.firstName.toLowerCase().includes(filter) || owner.lastName.toLowerCase().includes(filter));
    if(ownersContains.length > 0) {
      console.log('ownersContains', ownersContains)
      setOwnerListMatch(ownersContains.map(mapOwnerToSuggestionItem));
    } else {
        setOwnerListMatch([]);
      const ownersFuzzySearch = fuse.search(filter);
      console.log('ownersFuzzySearch', ownersFuzzySearch)
      if(ownersFuzzySearch.length > 0) {
        setOwnerListFuzzy(ownersFuzzySearch.map(p => p.item).map(mapOwnerToSuggestionItem));
      } else {
        setOwnerListFuzzy(ownerList.map(mapOwnerToSuggestionItem));
      }
    }
    setLoadingOwnerSuggestions(false);
  }

  const { data, loading, error } = useOrganizationOwner({ id });
  const { getUsersSuggestions } = useUserSuggestionsList();
  const { onLinkOrganizationOwner } = useLinkOrganizationOwner({
    organizationId: id,
    userId,
  });
  const { onUnlinkOrganizationOwner } = useUnlinkOrganizationOwner({
    organizationId: id,
  });

  useEffect(() => {
    if (!loading && data) {
      setInputValue(
        data.owner ? data.owner.firstName + ' ' + data.owner.lastName : '',
      );
    }
  }, [data, loading]);

  useEffect(() => {
    if (userId) {
      onLinkOrganizationOwner().then(() => {
        switchEditMode && switchEditMode();
        setUserId('');
      });
    }
  }, [userId]);

  return (
    <>
      <Autocomplete
        mode='full-width'
        debounce={0}
        editable={editMode}
        initialValue={inputValue}
        suggestionsMatch={ownerListMatch}
        suggestionsFuzzyMatch={ownerListFuzzy}
        onDoubleClick={() => {
          !editMode && switchEditMode && switchEditMode();
        }}
        onChange={(e: any) => {
          setUserId(e.value);
        }}
        loading={loadingOwnerSuggestions}
        onSearch={(filter: string) => searchOwners(filter)}
        onClearInput={() => {
          if (data?.owner) {
            onUnlinkOrganizationOwner();
          }
        }}
        placeholder='Owner...'
      />
    </>
  );
};
