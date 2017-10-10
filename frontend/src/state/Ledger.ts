import {} from '../model/transaction';
import { Transaction } from '../model';

class Ledger {
    constructor(changes: Partial<Ledger> = {}, previous?: Ledger) {
        Object.assign(this, previous, changes);

        if (previous && this.transactions !== previous.transactions) {
            this.transactions = Object.freeze(this.transactions.map(Object.freeze));
        }

        Object.freeze(this);
    }

    transactions: ReadonlyArray<Transaction> = Object.freeze([]);
}

export default Ledger;
