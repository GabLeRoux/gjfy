Vue.http.options.emulateJSON = true; // send as

new Vue({
        el: '#vueApp',
        delimiters: ['${', '}'],
        data: {
            // declare message with an empty value
            auth_token: '',
            secret: '',
            max_clicks: 1,
            postResults: {},
            requestInProgress: false
        },
        methods: {
            sendNewSecret: function () {
                this.requestInProgress = true;
                let data = JSON.stringify({
                    auth_token: this.auth_token,
                    secret: this.secret,
                    max_clicks: this.max_clicks
                });
                this.$http.post('/api/v1/new', data).then(function (response) {
                    // Success
                    this.postResults = response.data;
                    this.requestInProgress = false;
                }, function (response) {
                    // Error
                    this.postResults = response.data;
                    this.requestInProgress = false;
                });
            }
        }
    }
);