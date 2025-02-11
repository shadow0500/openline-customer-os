import React from 'react';
import styles from './issue-timeline-item.module.scss';
import Ticket from '@spaces/atoms/icons/Ticket';
import { TagsList } from '@spaces/atoms/tags/TagList';

// interface Props extends ContactWebAction {
//     contactName?: string
// }

export const IssueTimelineItem = ({
  createdAt,
  updatedAt,
  subject,
  status,
  priority,
  description,
  tags,
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
      </article>
    </div>
  );
};
