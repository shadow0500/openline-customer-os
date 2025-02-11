import * as React from 'react';
import { SVGProps } from 'react';
const SvgImages = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M18.5 3.75h-10A2.75 2.75 0 0 0 5.75 6.5v.25H5.5A2.75 2.75 0 0 0 2.75 9.5v8a2.75 2.75 0 0 0 2.75 2.75h10a2.75 2.75 0 0 0 2.75-2.75v-.25h.25a2.75 2.75 0 0 0 2.75-2.75v-8a2.75 2.75 0 0 0-2.75-2.75zM7.25 6.5A1.25 1.25 0 0 1 8.5 5.25h10a1.25 1.25 0 0 1 1.25 1.25v6.2l-2.27-1.91a.74.74 0 0 0-.542-.179.74.74 0 0 0-.508.259l-1.07 1.26-4-3.88a.7.7 0 0 0-.52-.25.75.75 0 0 0-.54.26l-3.05 3.63V6.5zm1.25 9.25a1.25 1.25 0 0 1-1.25-1.25v-.3l3.67-4.32 3.46 3.39-2.1 2.48H8.5zm8.25 1.75a1.25 1.25 0 0 1-1.25 1.25h-10a1.25 1.25 0 0 1-1.25-1.25v-8A1.25 1.25 0 0 1 5.5 8.25h.25v6.25a2.75 2.75 0 0 0 2.75 2.75h8.25v.25zm1.75-1.75h-4.25l2.84-3.34 2.63 2.23a1.23 1.23 0 0 1-1.22 1.11z'
      fill='currentColor'
    />
  </svg>
);
export default SvgImages;
