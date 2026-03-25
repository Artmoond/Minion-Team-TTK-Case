<script>
import NavPanel from './components/NavPanel.vue'
export default {
    name: 'App',
    components: {
        NavPanel,
    }
}
</script>


<template>
    <header class="w-full">
        <NavPanel/>
    </header>
    <main>
        <RouterView/>
    </main>
</template>