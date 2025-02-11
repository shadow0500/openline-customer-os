import * as React from 'react';
import { SVGProps } from 'react';
const SvgCloudUpload = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 20 16'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M15.5 14.25a.75.75 0 1 1 0-1.5c1.66 0 2.25-.83 2.25-3.18a3.57 3.57 0 0 0-3.25-3.25 3.34 3.34 0 0 0-1 .18.74.74 0 0 1-1-.49A5.25 5.25 0 0 0 2.25 7.57c0 3.44.76 5.18 2.25 5.18a.75.75 0 1 1 0 1.5C2 14.25.75 12 .75 7.57a6.75 6.75 0 0 1 13-2.68 4.4 4.4 0 0 1 .8-.07 5.07 5.07 0 0 1 4.75 4.75c-.05 1.28-.05 4.68-3.8 4.68z'
      fill='currentColor'
    />
    <path
      d='M12.83 11.65a.77.77 0 0 1-.53-.22L10 9.13l-2.3 2.3a.75.75 0 0 1-1.06-1.06l2.83-2.83a.74.74 0 0 1 1.06 0l2.83 2.83a.75.75 0 0 1 0 1.06.79.79 0 0 1-.53.22z'
      fill='currentColor'
    />
    <path
      d='M10 15.18a.75.75 0 0 1-.75-.75V8.07a.75.75 0 1 1 1.5 0v6.36a.75.75 0 0 1-.75.75z'
      fill='currentColor'
    />
  </svg>
);
export default SvgCloudUpload;
