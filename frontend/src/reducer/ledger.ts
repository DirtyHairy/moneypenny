import { MoneypennyAction } from '../action';
import { Actions } from '../action/ledger';
import State from '../state/Ledger';

function reducer(state = new State(), action: MoneypennyAction): State {
    switch (action.type) {
        case Actions.addTransaction:
            if (state.transactions.find(t => t.id === action.transaction.id)) {
                throw new Error(`transaction ${action.transaction.id} is already part of the collection`);
            }

            return new State(
                {
                    transactions: [
                        ...state.transactions.filter(t => t.id !== action.transaction.id),
                        action.transaction
                    ]
                },
                state
            );

        case Actions.removeTransaction:
            return new State({
                transactions: [...state.transactions.filter(t => t.id !== action.id)]
            });

        case Actions.replaceTransaction:
            if (!state.transactions.find(t => t.id === action.transaction.id)) {
                throw new Error(`transaction ${action.transaction.id} is not part of the collection`);
            }

            return new State(
                {
                    transactions: [
                        ...state.transactions.filter(t => t.id !== action.transaction.id),
                        action.transaction
                    ]
                },
                state
            );

        case Actions.replaceAllTransactions:
            return new State({
                transactions: action.transactions
            });

        default:
            return state;
    }
}

export default reducer;
