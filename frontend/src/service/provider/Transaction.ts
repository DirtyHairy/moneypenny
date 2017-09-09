import {Transaction, RequestResult} from '../../model';

export type LoadAllResult = RequestResult<Array<Transaction>>;

interface TransactionProvider {

    loadAll(): Promise<LoadAllResult>;

}

export default TransactionProvider;
