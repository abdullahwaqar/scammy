<template>
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
            <a class="navbar-item" href="/">
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
</template>

<script>
export default {
    name: 'Navbar'
}
</script>