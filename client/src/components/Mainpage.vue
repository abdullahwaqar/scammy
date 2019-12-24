<template>
<section>

    <section class="hero is-light is-fullheight">
    <!-- Hero head: will stick at the top -->
    <div class="hero-head">
        <header class="navbar">
        <div class="container">
            <div class="navbar-brand">
            <a class="navbar-item" href="/">
                <h1 class="logo-font is-size-4">Scammy</h1>
            </a>
            <span class="navbar-burger burger" data-target="navbarMenuHeroC">
                <span></span>
                <span></span>
                <span></span>
            </span>
            </div>
            <div id="navbarMenuHeroC" class="navbar-menu">
            <div class="navbar-end">
                <a class="navbar-item is-active" href="/">
                Home
                </a>
                <a class="navbar-item" href="/report">
                Report
                </a>
                <a class="navbar-item">
                About
                </a>
            </div>
            </div>
        </div>
        </header>
    </div>

    <!-- Hero content: will be in the middle -->
    <div class="hero-body">
        <div class="container has-text-centered">
        <h1 class="title logo-font">
            Scammy
        </h1>
        <hr>
        <h2 class="subtitle">
            Received a fishy email? Can it be a scammer? Check if the same email was sent to others or <u><a href="/report">report</a></u> the email.
        </h2>
        <search-bar></search-bar>
        </div>
    </div>
    </section>
</section>
</template>

<script>
import LevelItem from '@/components/LevelItem.vue';
import Search from '@/components/Search.vue';

export default {
    name: 'MainPage',
    components: {
        'levelitem': LevelItem,
        'search-bar': Search
    }
}
</script>

<style scoped>
.logo-font {
    font-family: 'Tomorrow', sans-serif;
}
</style>
