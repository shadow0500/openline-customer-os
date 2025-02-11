import React, { FC, useRef, useState } from 'react';
import { MeetingParticipant } from '../../../../../graphQL/__generated__/generated';
import styles from './attendee-autocomplete.module.scss';
import { IconButton } from '@spaces/atoms/icon-button/IconButton';
import classNames from 'classnames';
import { useDetectClickOutside } from '@spaces/hooks/useDetectClickOutside';
import { ContactAvatar } from '../../contact-avatar/ContactAvatar';
import { getAttendeeDataFromParticipant } from '../utils';

interface PreviewAttendeesProps {
  selectedAttendees: Array<MeetingParticipant>;
  hiddenAttendeesNumber: number;
}

export const PreviewAttendees: FC<PreviewAttendeesProps> = ({
  selectedAttendees = [],
  hiddenAttendeesNumber = 0,
}) => {
  const [dropdownOpen, setDropdownOpen] = useState<boolean>(false);

  const attendeeAutocompleteWrapperRef = useRef(null);

  useDetectClickOutside(attendeeAutocompleteWrapperRef, () => {
    if (!dropdownOpen) return;
    setDropdownOpen(false);
  });

  return (
    <div ref={attendeeAutocompleteWrapperRef} style={{ position: 'relative' }}>
      <IconButton
        mode='secondary'
        label='See all attendees'
        onClick={() => setDropdownOpen(!dropdownOpen)}
        icon={<span style={{ fontSize: 12 }}>+{hiddenAttendeesNumber}</span>}
        size='xxxs'
        style={{
          width: 30,
          height: 30,
          zIndex: 3,
          // left: -60,
          position: 'relative',
        }}
      />

      {dropdownOpen && (
        <div
          className={classNames(
            styles.attendeeAutocompleteWrapper,
            styles.right,
          )}
        >
          {selectedAttendees.map((attendeeData) => {
            const attendee = getAttendeeDataFromParticipant(attendeeData);
            return (
              <li
                key={`attendee-suggestion-${attendee.id}`}
                className={classNames(styles.suggestionItem, styles.selected)}
              >
                <ContactAvatar size={20} contactId={attendee.id} showName />
              </li>
            );
          })}
        </div>
      )}
    </div>
  );
};
