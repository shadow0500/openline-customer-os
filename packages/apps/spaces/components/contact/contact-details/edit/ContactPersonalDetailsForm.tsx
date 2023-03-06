import React, { FormEvent, useMemo, useRef } from 'react';
import Image from 'next/image';
import { Button, User, UserPlus } from '../../../ui-kit/atoms';
import styles from '../contact-details.module.scss';
import { ContactTags } from '../../contact-tags';
import { ControlledInput } from '../../../ui-kit/atoms/input/ControlledInput';
import { Controller, useForm } from 'react-hook-form';
import { PersonTitle } from '../../../../graphQL/__generated__/generated';
import { OverlayPanel } from '../../../ui-kit/atoms/overlay-panel';
import { Dropdown } from 'primereact/dropdown';
import { capitalizeFirstLetter } from '../../../../utils';

export const ContactPersonalDetailsEditForm = ({
  data,
  onSubmit,
  mode = 'EDIT',
}: {
  data: any;
  onSubmit: any;
  mode?: 'CREATE' | 'EDIT';
}) => {
  const titleSelectorRef = useRef(null);

  const { control, getValues } = useForm({
    defaultValues: {
      firstName: data?.firstName || '',
      id: data.id,
      label: data?.label || '',
      lastName: data?.lastName || '',
      ownerId: data?.ownerId || '',
      title: data?.title || PersonTitle.Mr,
    },
  });

  const handleSubmit = (e: FormEvent<HTMLButtonElement>, values: any) => {
    e.stopPropagation();
    e.preventDefault();
    onSubmit(values);
  };

  const titleOptions = useMemo(
    () =>
      Object.values(PersonTitle).map((labelOption) => ({
        value: labelOption,
        label: capitalizeFirstLetter(labelOption),
      })),
    [],
  );

  return (
    <form className={styles.header}>
      {mode === 'EDIT' && (
        <div className={styles.photo}>
          {data?.photo ? (
            <Image src={data?.photo} alt={''} height={40} width={40} />
          ) : (
            <User />
          )}
        </div>
      )}

      <div className={styles.name}>
        {mode === 'CREATE' && (
          <div style={{ display: 'flex', justifyContent: 'center' }}>
            <div className={styles.photo}>
              <UserPlus />
            </div>
          </div>
        )}

        <div
          className={styles.formFields}
          style={{ display: 'flex', flexDirection: 'column', width: '100%' }}
        >
          <label className={styles.xyz} htmlFor='title'>
            Title
          </label>
          <Controller
            name='title'
            control={control}
            render={({ field }) => (
              <Dropdown
                id={field.name}
                value={field.value}
                onChange={(e) => field.onChange(e.value)}
                options={titleOptions}
                optionValue='value'
                optionLabel='label'
                className={styles.titleSelector}
              />
            )}
          />
          <OverlayPanel ref={titleSelectorRef} model={titleOptions} />

          <ControlledInput
            required={true}
            inputSize='xxxs'
            control={control}
            name='firstName'
            placeholder=''
            label='First name'
          />
          <ControlledInput
            required={true}
            inputSize='xxxs'
            control={control}
            name='lastName'
            placeholder=''
            label='Last name'
          />
        </div>

        {mode === 'EDIT' && <ContactTags id={data.id} mode={'EDIT'} />}
        <div
          style={{
            display: 'flex',
            justifyContent: 'flex-end',
            marginTop: '8px',
          }}
        >
          {}
          <Button
            mode='primary'
            type='button'
            onClick={(e) => handleSubmit(e, getValues())}
          >
            {mode === 'CREATE' ? 'Create contact' : 'Save'}
          </Button>
        </div>
      </div>
    </form>
  );
};
