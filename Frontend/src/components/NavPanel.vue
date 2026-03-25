<script>
import Button from './Button.vue'

export default {
    name: 'NavPanel',
    components: {
        Button,
    },
    methods: {
        navigateToLogin() {
            this.$router.push('/')
        },
        navigateToAdmin() {
            this.$router.push('/admin')
        },
        navigateToHome() {
            this.$router.push('/player')
        },
        navigateToHost() {
            this.$router.push('/host')
        },
    }
}

</script>
<template>
    <section class="flex flex-row gap-2">
        <Button
            text="Авторизация"
            @click="navigateToLogin"
        />
        <Button
            text="Администрирование"    
            @click="navigateToAdmin"
        />
        <Button
            text="Главная страница"
            @click="navigateToHome"
        />
        <Button
            text="Раздел ведущего"
            @click="navigateToHost"
        />
    </section>
</template>
