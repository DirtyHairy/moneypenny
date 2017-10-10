import { Action } from 'redux';

import { Transaction } from '../model';

export enum Actions {
    addTransaction = 'ledger/add-transaction',
    removeTransaction = 'ledger/remove-transaction',
    replaceTransaction = 'ledger/replace-transaction',
    replaceAllTransactions = 'ledger/replace-all-transactions'
}

export type LedgerAction =
    | AddTransactionAction
    | RemoveTransactionAction
    | ReplaceTransactionAction
    | ReplaceAllTransactionsActions;

export interface AddTransactionAction extends Action {
    type: Actions.addTransaction;
    transaction: Transaction;
}

export function addTransaction(transaction: Transaction): AddTransactionAction {
    return {
        type: Actions.addTransaction,
        transaction
    };
}

export interface RemoveTransactionAction extends Action {
    type: Actions.removeTransaction;
    id: number;
}

export function removeTransaction(id: number): RemoveTransactionAction {
    return {
        type: Actions.removeTransaction,
        id
    };
}

export interface ReplaceTransactionAction extends Action {
    type: Actions.replaceTransaction;
    transaction: Transaction;
}

export function replaceTransaction(transaction: Transaction): ReplaceTransactionAction {
    return {
        type: Actions.replaceTransaction,
        transaction
    };
}

export interface ReplaceAllTransactionsActions extends Action {
    type: Actions.replaceAllTransactions;
    transactions: Array<Transaction>;
}

export function replaceAllTransactions(transactions: Array<Transaction>): ReplaceAllTransactionsActions {
    return {
        type: Actions.replaceAllTransactions,
        transactions
    };
}
