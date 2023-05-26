import React, { useEffect, useRef, useState } from 'react';

import { SuggestionList } from './suggestion-list/SuggestionList';
import { useGCLI } from './context/GCLIContext';
import Search from '@spaces/atoms/icons/Search';
import { DebouncedInput } from '@spaces/atoms/input/DebouncedInput';
import styles from './GCLIInput.module.scss';
import { uuid4 } from '@sentry/utils';

// TODO
// Filtering:
// 1. executing filter action on enter is liniting next filter options to specific context eg. organisation
// 2. executing filter action on comma is grouping objects, no action executed (allows to execute action on all of those objects)
// objects should be initially of the same type

// TODO
//   add result context change mechanism

export const GCLIInput = () => {
  // TODO action simulation to be removed!

  const {
    label,
    icon,
    inputPlaceholder,
    existingTerms,
    loadSuggestions,
    loadingSuggestions,
    suggestionsLoaded,
    onItemsChange,
    selectedTermFormat,
  } = useGCLI();

  // todo use input value to create fill in effect on navigate through results by keyboard ??? do we even need that? is this a proper use case
  const [selectedValues, setSelectedValues] = useState(existingTerms ?? []);
  const [locationTerms, setLocationTerms] = useState([] as any[]);
  const [searchTerms, setSearchTerms] = useState([] as any[]);
  const [suggestionListKeyDown, setSuggestionListKeyDown] =
    useState<boolean>(false);

  useEffect(() => {
    setLocationTerms(
      selectedValues.filter((item: any) => item.type === 'STATE'),
    );
    setSearchTerms(
      selectedValues.filter((item: any) => item.type === 'GENERIC'),
    );
  }, [selectedValues]);

  const [searchQuery, setSearchQuery] = useState('');

  const [suggestions, setSuggestions] = useState<Array<any>>([]);

  const inputRef = useRef<HTMLInputElement>(null);

  const [dropdownOpen, setDropdownOpen] = useState(false);
  const dropdownRef = React.useRef(null);

  useEffect(() => {
    if (!loadingSuggestions && suggestionsLoaded) {
      setSuggestions(suggestionsLoaded);
    }
  }, [loadingSuggestions, suggestionsLoaded]);

  // HANDLERS FOR GENERAL ACTIONS
  const handleSearchResultSelect = (item: any | undefined) => {
    if (item === undefined) {
      setDropdownOpen(false);
      inputRef.current?.focus();
      return;
    }
    setDropdownOpen(false);
    const items = [...selectedValues, item];
    setSelectedValues(items);
    onItemsChange(items);
    setSearchQuery('');
    setSuggestions([]);
    inputRef?.current?.focus();
  };

  // HANDLERS FOR GENERAL ACTIONS
  const handleAsSimpleSearch = () => {
    if (!searchQuery) return;
    handleSearchResultSelect({
      id: uuid4(),
      type: 'GENERIC',
      display: searchQuery,
    });
  };
  // END HANDLERS FOR GENERAL ACTIONS

  const handleInputKeyDown = (event: any) => {
    const { key, currentTarget, target } = event;
    switch (key) {
      case 'Enter':
        handleAsSimpleSearch();
        break;
      case 'ArrowDown':
        setSuggestionListKeyDown(!suggestionListKeyDown);
        break;
      case 'Escape':
        setDropdownOpen(false);
    }
  };

  const handleInputChange = (event: any) => {
    if (!event.target.value) {
      setDropdownOpen(false);
      setSuggestions([]);
      return;
    }
    setSearchQuery(event.target.value);
    inputRef.current?.focus();
    setDropdownOpen(true);

    loadSuggestions(event.target.value);
  };

  return (
    <div className={styles.gcli_wrapper}>
      <div className={styles.input_wrapper}>
        <div className={styles.input_label_icon}>
          {icon && <div className={styles.input_icon}>{icon}</div>}

          <div className={styles.input_label}>{label}</div>
        </div>

        <div className={styles.selected_terms_wrapper}>
          {locationTerms.length > 0 && <span>in:&nbsp;</span>}
          {locationTerms.map((e, index) => {
            return (
              <div
                className={styles.selected_term}
                key={index}
                onClick={() => {
                  const filters = selectedValues.filter(
                    (term, i) =>
                      term.type !== e.type ||
                      (term.type === e.type && term.id !== e.id),
                  );
                  setSelectedValues(filters);
                  onItemsChange(filters);
                  inputRef?.current?.focus();
                }}
              >
                <div className={styles.selected_term_text}>
                  {selectedTermFormat ? selectedTermFormat(e) : e.display}
                  {index < locationTerms.length - 1 ? ',' : ''}
                </div>
              </div>
            );
          })}
          {searchTerms.length > 0 && <span>contains:&nbsp;</span>}
          {searchTerms.map((e, index) => (
            <div
              className={styles.selected_term}
              key={index}
              onClick={() => {
                const filters = selectedValues.filter(
                  (term, i) =>
                    term.type !== e.type ||
                    (term.type === e.type && term.id !== e.id),
                );
                setSelectedValues(filters);
                onItemsChange(filters);
                inputRef?.current?.focus();
              }}
            >
              <div className={styles.selected_term_text}>
                {selectedTermFormat ? selectedTermFormat(e) : e.display}
                {index < searchTerms.length - 1 ? ',' : ''}
              </div>
            </div>
          ))}
        </div>

        <DebouncedInput
          inputRef={inputRef}
          placeholder={inputPlaceholder}
          className={styles.gcli_input_search}
          minLength={1}
          value={searchQuery}
          onChange={handleInputChange}
          onKeyDown={handleInputKeyDown}
          debounceTimeout={50}
        />

        <div className={styles.input_actions}>
          {!loadingSuggestions && searchQuery !== '' && (
            <button
              className={styles.search_button}
              onClick={handleAsSimpleSearch}
            >
              <Search height={16} style={{ marginRight: '10px' }} />
              Search
            </button>
          )}
        </div>
      </div>
      {/* END SELECTED OPTIONS */}

      {dropdownOpen && searchQuery !== '' && (
        <SuggestionList
          onSearchResultSelect={handleSearchResultSelect}
          loadingSuggestions={loadingSuggestions}
          suggestions={suggestions}
          suggestionListKeyDown={suggestionListKeyDown}
          onIndexChanged={(index: number | null) => {
            if (index === null) {
              inputRef?.current?.focus();
            }
          }}
        />
      )}
    </div>
  );
};
