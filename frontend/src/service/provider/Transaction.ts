import {Transaction, RequestResult} from '../../model';

export type LoadAllResult = Promise<RequestResult<Array<Transaction>>>;

interface TransactionProvider {

    loadAll(): LoadAllResult;

}

export default TransactionProvider;
