<script>
import Plyr from 'plyr';
import 'plyr/dist/plyr.css';
import Input from '../components/Input.vue';
import Button from '../components/Button.vue';

export default {
    name: 'PagePlayer',
    components: {
        Input,
        Button
    },
    mounted() {
        const player = new Plyr('#audio-player');
    }
}
</script>


<template>
    <section class="flex flex-row gap-2 w-full p-4">
        <div class="player-container w-full">
            <div class="player-card p-2 border border-[var(--gray)]">
                <audio id="audio-player" controls crossorigin>
                    <source src="http://localhost:8081/stream/StartAudio" type="audio/mpeg">
                    Ваш браузер не поддерживает аудио
                </audio>
            </div>
        </div>
        <div>
            <span class="w-full flex flex-row gap-2">
                <Input
                    class="w-full"
                    placeholder="Сообщение"
                    allowed-alphabets="ru, en, symbols, numeric"
                />
                <Button text="Отправить" class="h-fit"/>         
            </span>
            <div class="border border-[var(--gray)] p-2 h-fit">
                <p class="text-sm">тестовое сообщение</p>
                <p class="text-sm">тестовое сообщение</p>
                <p class="text-sm">тестовое сообщение</p>        
            </div>
        </div>
    </section>
</template>