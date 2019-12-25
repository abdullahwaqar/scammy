<template>
    <section class="section">
        <div v-if="error" class="notification is-danger">
        <button @click="error = ''" class="delete"></button>
        {{ error }}
        </div>
        <div class="columns">
            <div class="column is-half is-offset-one-quarter">
            <div class="control">
                <input v-model="searchTerm" class="input is-medium is-rounded" type="email" placeholder="Enter Email Address">
            </div>
            </div>
        </div>
        <br>
        <button @click="search" class="button is-medium is-dark is-capitalized">Check Email?</button>
    </section>
</template>

<script>
import { mapActions } from 'vuex';

export default {
    name: 'Search',
    data: function() {
        return {
            searchTerm: '',
            error: ''
        }
    },
    methods: {
        ...mapActions({ getEmail: 'getEmailDetails'}),
        validateEmail(email) {
            var re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return re.test(email);
        },
        search() {
            if (!this.validateEmail(this.searchTerm)) {
                this.error = 'This not a valid email.'
                return;
            }
            this.getEmail(this.searchTerm);
        }
    },
    watch: {
        searchTerm: function() {
            this.error = '';
        }
    }
}
</script>