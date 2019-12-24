<template>
    <section class="section">
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
            searchTerm: ''
        }
    },
    methods: {
        search() {}
    }
}
</script>