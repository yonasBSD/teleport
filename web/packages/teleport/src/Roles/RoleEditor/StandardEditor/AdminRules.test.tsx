/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { act } from '@testing-library/react';
import selectEvent from 'react-select-event';

import { render, screen, userEvent } from 'design/utils/testing';
import { Validator } from 'shared/components/Validation';

import { ResourceKind } from 'teleport/services/resources';

import { AdminRules } from './AdminRules';
import { RuleModel } from './standardmodel';
import { StatefulSectionWithDispatch } from './StatefulSection';
import { StandardModelDispatcher } from './useStandardModel';
import { AdminRuleValidationResult } from './validation';

describe('AdminRules', () => {
  const setup = () => {
    const modelRef = jest.fn();
    let validator: Validator;
    let dispatch: StandardModelDispatcher;
    render(
      <StatefulSectionWithDispatch<RuleModel[], AdminRuleValidationResult[]>
        selector={m => m.roleModel.rules}
        validationSelector={m => m.validationResult.rules}
        component={AdminRules}
        validatorRef={v => {
          validator = v;
        }}
        modelRef={modelRef}
        dispatchRef={d => {
          dispatch = d;
        }}
      />
    );
    return { user: userEvent.setup(), modelRef, dispatch, validator };
  };

  test('editing', async () => {
    const { user, modelRef } = setup();
    await user.click(screen.getByRole('button', { name: 'Add New' }));
    await selectEvent.select(screen.getByLabelText('Teleport Resources *'), [
      'db',
      'node',
    ]);
    await selectEvent.select(screen.getByLabelText('Permissions *'), [
      'list',
      'read',
    ]);
    await user.type(screen.getByLabelText('Filter'), 'some-filter');
    expect(modelRef).toHaveBeenLastCalledWith([
      {
        id: expect.any(String),
        resources: [
          { label: ResourceKind.Database, value: 'db' },
          { label: ResourceKind.Node, value: 'node' },
        ],
        verbs: [
          { label: 'list', value: 'list' },
          { label: 'read', value: 'read' },
        ],
        where: 'some-filter',
        hideValidationErrors: true,
      },
    ] as RuleModel[]);
  });

  test('validation', async () => {
    const { user, validator, dispatch } = setup();
    await user.click(screen.getByRole('button', { name: 'Add New' }));
    act(() => validator.validate());

    // Validation hidden
    expect(
      screen.queryByText('At least one resource kind is required')
    ).not.toBeInTheDocument();
    expect(
      screen.queryByText('At least one permission is required')
    ).not.toBeInTheDocument();

    // Validation visible
    act(() => dispatch({ type: 'show-validation-errors' }));
    expect(
      screen.getByText('At least one resource kind is required')
    ).toBeInTheDocument();
    expect(
      screen.getByText('At least one permission is required')
    ).toBeInTheDocument();
  });
});
