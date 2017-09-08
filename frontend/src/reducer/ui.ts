import {MoneypennyAction} from '../action';
import {Actions} from '../action/ui';
import State from '../state/ui';

function reducer(state = new State(), action: MoneypennyAction): State {
    switch (action.type) {

        case Actions.setLoading:
            return new State({loading: action.loading}, state);

        default:
            return state;
    }
}

export default reducer;
