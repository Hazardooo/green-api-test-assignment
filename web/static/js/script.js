class GreenAPIClient {
    constructor(apiBase = '/api') {
        this.apiBase = apiBase;
        this.elements = this.cacheElements();
        this.init();
    }

    cacheElements() {
        const ids = [
            'idInstance', 'apiToken', 'chatId', 'messageText',
            'fileChatId', 'fileUrl', 'fileName', 'fileCaption',
            'output', 'statusBar'
        ];

        const elements = {};
        ids.forEach(id => {
            elements[id] = document.getElementById(id);
        });
        return elements;
    }

    init() {
        this.loadFromStorage();
        this.bindEvents();
    }

    loadFromStorage() {
        const storageKeys = [
            'idInstance', 'apiToken', 'chatId', 'messageText',
            'fileChatId', 'fileUrl', 'fileName', 'fileCaption'
        ];

        storageKeys.forEach(key => {
            const saved = localStorage.getItem(key);
            if (saved && this.elements[key]) {
                this.elements[key].value = saved;
            }
        });
    }

    bindEvents() {
        const storageKeys = [
            'idInstance', 'apiToken', 'chatId', 'messageText',
            'fileChatId', 'fileUrl', 'fileName', 'fileCaption'
        ];

        storageKeys.forEach(key => {
            const el = this.elements[key];
            if (el) {
                el.addEventListener('input', () => {
                    localStorage.setItem(key, el.value);
                });
            }
        });
    }

    toggleParams(type) {
        const el = document.getElementById(type + 'Params');
        if (el) {
            el.style.display = el.style.display === 'none' ? 'block' : 'none';
        }
    }

    clearOutput() {
        if (this.elements.output) {
            this.elements.output.value = '';
        }
        this.updateStatus('Готов к работе', 'text-muted');
    }

    updateStatus(text, colorClass) {
        const bar = this.elements.statusBar;
        if (bar) {
            bar.textContent = text;
            bar.className = 'mt-2 small fw-bold ' + colorClass;
        }
    }

    getCredentials() {
        return {
            idInstance: this.elements.idInstance?.value.trim() || '',
            apiToken: this.elements.apiToken?.value.trim() || ''
        };
    }

    validateCredentials() {
        const { idInstance, apiToken } = this.getCredentials();
        if (!idInstance || !apiToken) {
            alert('Введите idInstance и ApiTokenInstance!');
            return false;
        }
        return true;
    }

    buildRequestBody(method) {
        const { idInstance, apiToken } = this.getCredentials();
        let body = { idInstance, apiToken };

        switch(method) {
            case 'message':
                body.chatId = this.elements.chatId?.value;
                body.message = this.elements.messageText?.value;
                break;
            case 'file':
                body.chatId = this.elements.fileChatId?.value;
                body.urlFile = this.elements.fileUrl?.value;
                body.fileName = this.elements.fileName?.value;
                body.caption = this.elements.fileCaption?.value;
                break;
        }

        return body;
    }

    async callApi(method, btn) {
        if (!this.validateCredentials()) return;

        this.setLoading(btn, true);
        this.updateStatus('Выполняется запрос...', 'text-primary');

        try {
            const url = `${this.apiBase}/${method}`;
            const body = this.buildRequestBody(method);

            const response = await fetch(url, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(body)
            });

            if (!response.ok) {
                const text = await response.text();
                throw new Error(`HTTP ${response.status}: ${text || response.statusText}`);
            }

            const data = await response.json();
            this.displayResult(data, method);

        } catch (error) {
            this.displayError(error);
        } finally {
            this.setLoading(btn, false);
        }
    }

    displayResult(data, method) {
        if (this.elements.output) {
            this.elements.output.value = JSON.stringify(data, null, 2);
        }
        this.updateStatus(
            `${method} выполнен успешно (${new Date().toLocaleTimeString()})`,
            'text-success'
        );
    }

    displayError(error) {
        if (this.elements.output) {
            this.elements.output.value = `Ошибка: ${error.message}`;
        }
        this.updateStatus('Ошибка выполнения', 'text-danger');
        console.error(error);
    }

    setLoading(btn, isLoading) {
        if (isLoading) {
            btn.classList.add('loading');
            btn.disabled = true;
        } else {
            btn.classList.remove('loading');
            btn.disabled = false;
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    window.apiClient = new GreenAPIClient();
});

function toggleParams(type) {
    window.apiClient.toggleParams(type);
}

function clearOutput() {
    window.apiClient.clearOutput();
}

function callApi(method, btn) {
    window.apiClient.callApi(method, btn);
}