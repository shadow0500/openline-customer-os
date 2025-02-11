import React, { ButtonHTMLAttributes, FC, ReactNode } from 'react';
import styles from './button.module.scss';
import classNames from 'classnames';

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  icon?: ReactNode;
  ariaLabel?: string;
  children?: React.ReactNode;
  mode?:
    | 'default'
    | 'primary'
    | 'secondary'
    | 'danger'
    | 'link'
    | 'dangerLink'
    | 'text';
}

export const Button: FC<Props> = ({
  icon,
  onClick,
  children,
  mode = 'default',
  ...rest
}) => {
  return (
    <button
      {...rest}
      onClick={onClick}
      className={classNames(styles.button, styles[mode], rest.className)}
    >
      <>
        {icon && icon}
        {children}
      </>
    </button>
  );
};
