<script>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";

export default {
    components: {Input, Button},
    methods: {
        toggleLogIn() {
            this.isLogIn = !this.isLogIn;
        }
    },
    data() {
        return {
            isLogIn: true,
        }
    }
}
</script>

<template>
    <section
        class="flex flex-col items-center gap-4"
        v-if="isLogIn"
    >
        <p class="text-xl">Вход</p>
        <div class="w-full flex flex-col gap-1">
            <Input
                class="w-full"
                placeholder="Логин"
                allowed-alphabets="en"
                meta="login"
            />
            <Input
                class="w-full"
                placeholder="Пароль"
                allowed-alphabets="en, symbols, numeric"
                meta="password"
            />
        </div>
        <Button text="Войти" class="w-full"/>
        <a class="underline cursor-pointer" @click='toggleLogIn' >Регистрация</a>
    </section>

    <section
        class="flex flex-col items-center gap-4"
        v-if="!isLogIn"
    >
        <p class="text-xl">Регистрация</p>
        <div class="w-full flex flex-col gap-1">
            <Input
                placeholder="Логин"
                allowed-alphabets="en"
                meta="login"
            />
            <Input
                placeholder="ФИО"
                allowed-alphabets="ru"
            />
            <Input
                placeholder="Пароль"
                allowed-alphabets="en, symbols, numeric"
                meta="password"
            />
            <Input
                placeholder="Пароль"
                allowed-alphabets="en, symbols, numeric"
                meta="password"
            />
        </div>
        <Button
            text="Войти"
            class="w-full"
        />
        <a class="underline cursor-pointer" @click='toggleLogIn' >Вход</a>
    </section>
</template>

<style scoped>
section{
    width: 400px;
    border: black solid 1px;
    padding: 12px;
    border-radius: 12px;
}
</style>