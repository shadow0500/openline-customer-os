import React from 'react';

import styles from './ticket-timeline-item.module.scss';
import Ticket from '../../atoms/icons/Ticket';
import { TagsList } from '../../atoms';
import sanitizeHtml from 'sanitize-html';
import format from 'date-fns/format';
import { DateTimeUtils } from '../../../../utils';

// interface Props extends ContactWebAction {
//     contactName?: string
// }

export const TicketTimelineItem = ({
  createdAt,
  updatedAt,
  subject,
  status,
  priority,
  description,
  tags,
  notes,
  ...rest
}: any): JSX.Element => {
  return (
    <div className={styles.x}>
      <article className={`${styles.ticketContainer}`}>
        <div className={`${styles.ticketHeader}`}>
          <div className={`${styles.ticketHeaderSubject}`}>
            <Ticket className={`${styles.ticketHeaderPicture}`} />
            {subject}
          </div>
          <div className={`${styles.ticketHeaderStatus}`}>{status}</div>
        </div>

        {tags && tags.length > 0 && (
          <div className={`${styles.tags}`}>
            <TagsList tags={tags ?? []} readOnly={true} />
          </div>
        )}

        <div>{description}</div>

        {description && notes && notes.length > 0 && (
          <div className={`${styles.contentNotesSeparator}`}></div>
        )}

        {notes &&
          notes.length > 0 &&
          notes.map((note: any) => {
            return (
              <div key={note.id}>
                <div className={`${styles.noteActivity}`}>
                  Activity on {DateTimeUtils.format(new Date(note.createdAt))}
                </div>
                <div
                  dangerouslySetInnerHTML={{ __html: sanitizeHtml(note.html) }}
                ></div>
              </div>
            );
          })}
      </article>
    </div>
  );
};
