import * as React from 'react';
import { SVGProps } from 'react';
const SvgPencil = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M4.21 20.52a.73.73 0 0 1-.52-.21.75.75 0 0 1-.22-.6l.31-3.84A.73.73 0 0 1 4 15.4L15.06 4.34a3.19 3.19 0 0 1 2.28-.86 3.3 3.3 0 0 1 2.25.91 3.31 3.31 0 0 1 .11 4.5L8.63 20a.77.77 0 0 1-.46.22l-3.89.35-.07-.05zm1-4.26L5 19l2.74-.25 10.9-10.92a1.72 1.72 0 0 0 .243-1.842A1.72 1.72 0 0 0 17.31 5a1.61 1.61 0 0 0-1.19.42L5.21 16.26z'
      fill='currentColor'
    />
  </svg>
);
export default SvgPencil;
