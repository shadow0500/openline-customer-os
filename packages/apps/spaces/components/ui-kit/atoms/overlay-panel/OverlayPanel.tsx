import React, { ForwardedRef, forwardRef, LegacyRef } from 'react';
import { Menu, MenuProps } from 'primereact/menu';
import {
  OverlayPanel as PrimereactOverlayPanel,
  OverlayPanelProps,
} from 'primereact/overlaypanel';

export const OverlayPanel = forwardRef(function MyInput(
  { model, ...rest }: MenuProps & OverlayPanelProps,
  ref: ForwardedRef<LegacyRef<PrimereactOverlayPanel> | null>,
) {
  return (
    <PrimereactOverlayPanel
      // @ts-expect-error primereact doesn't count in possibility of forwarding ref in their types
      ref={ref}
      dismissable
      {...rest}
      style={{ background: '#fff', ...rest.style }}
    >
      {rest.children && rest.children}

      {model && <Menu model={model} />}
    </PrimereactOverlayPanel>
  );
});
