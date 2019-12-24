<template>
    <footer class="footer">
        <nav class="level">
        <div class="level-item has-text-centered">
            <div>
            </div>
        </div>
        <div class="level-item has-text-centered">
            <div>
                <div class="field has-addons">
                <div class="content has-text-centered">
                <p>
                    <strong>
                        <a href="/">Scammy</a>
                    </strong>
                    Made with ❤ by <a href="https://github.com/abdullahwaqar">Abdullah</a>.
                </p>
                </div>
                    </div>
            </div>
        </div>
        <div class="level-item has-text-centered">
            <div>
                Status: <span :class="['fas', 'fa-dot-circle', status ? 'has-text-success' : 'has-text-danger']"></span>
            </div>
        </div>
        <div class="level-item has-text-centered">
            <div>
            </div>
        </div>
        </nav>
    </footer>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
    name: 'Footer',
    computed: {
        ...mapGetters({ status: 'getStatus' })
    },
    mounted: function() {
        this.ping();
    },
    methods: {
        ...mapActions({ ping: 'pingServer' })
    }
}
</script>
