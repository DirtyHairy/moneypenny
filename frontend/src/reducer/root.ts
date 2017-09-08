import {Action} from 'redux';
import State from '../state/state';

import {MoneypennyAction} from '../action';
import ledgerReducer from './ledger';
import uiReducer from './ui';

function reducer(state: State = {}, action: Action): State {
    return {
        ledger: ledgerReducer(state.ledger, action as MoneypennyAction),
        ui: uiReducer(state.ui, action as MoneypennyAction)
    };
}

export default reducer;
