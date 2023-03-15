import type { NextPage } from 'next';
import * as React from 'react';
import { FC, useEffect, useState } from "react";

interface SuggestionListProps {
  currentValue: string;
  getSuggestions: Function;
  setCurrentValue: Function;
}

interface _Suggestion {
  id: string;
  display: string;
  value: string;
}
export const SuggestionList: React.FC<any> = (props: SuggestionListProps) => {
  const [suggestions, setSuggestions] = useState(Array<_Suggestion>);
  const [showSuggestions, setShowSuggestions] = useState(false);

  useEffect(() => {
    if (props.currentValue.length > 0) {
      props.getSuggestions(props.currentValue, (s: _Suggestion[]) => {
        console.log('Setting suggestions: ' + JSON.stringify(s));
        setSuggestions(s);
        setShowSuggestions(true);
      });
    } else {
      setShowSuggestions(false);
    }
  }, [props.currentValue]);

  const selectSuggestion = (suggestion: string) => {
    props.setCurrentValue(suggestion);
    setShowSuggestions(false);
  };

  return (
    <div className='suggestion-list'>
      {showSuggestions &&
        suggestions &&
        suggestions.length > 0 &&
        suggestions.map((suggestion, index) => {
          return (
            <div
              className='cta'
              key={suggestion.id}
              onClick={() => selectSuggestion(suggestion.value)}
            >
              {suggestion.display} ({suggestion.value})
            </div>
          );
        })}
      {showSuggestions && suggestions && suggestions.length == 0 && (
        <div>No suggestions</div>
      )}
    </div>
  );
};
