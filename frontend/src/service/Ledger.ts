import {} from '../state/ledger';
interface LedgerService {
    start(): Promise<void>;
}

export default LedgerService;
