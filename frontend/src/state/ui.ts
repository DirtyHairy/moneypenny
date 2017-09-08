class Ui {

    constructor(changes: Partial<Ui> = {}, previous?: Ui) {
        Object.assign(this, previous, changes);
    }

    loading: boolean = false;

}

export default Ui;
