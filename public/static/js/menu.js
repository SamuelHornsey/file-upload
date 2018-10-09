class Menu extends HTMLElement {
    constructor () {
        super();

        this.confirm = document.querySelector('.js-confirm');

        this._hide = this._hide.bind(this);
        this._delete = this._delete.bind(this);

        window.addEventListener('click', this._hide);
        this.querySelector('.js-delete').addEventListener('click', this._delete);
        this.querySelector('.js-download').addEventListener('click', this._download);
    }

    _hide () {
        this.classList.remove('show');
    }

    _delete () {
        console.log('Delete')
        this.confirm.classList.add('show');
    }

    _download () {
        console.log('Download')
    }

    show (event) {
        this.style.transform = '';
        let bound = this.getBoundingClientRect();
        this.style.transform = `translateX(${event.clientX - bound.x}px) translateY(${event.clientY - bound.y}px)`;
        this.classList.add('show');
    }
}

customElements.define('app-menu', Menu);