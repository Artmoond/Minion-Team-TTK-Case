<script setup>
import { computed } from 'vue';

const props = defineProps({
    Value:         { type: String, default: '' },
    placeholder:        String,
    meta:               { type: String, default: 'text' },
    allowedAlphabets:   { type: Array, default: () => [] }
});

const emit = defineEmits(['update:modelValue']);

const inputType = computed(() => {
    return props.meta === 'password' ? 'password' : 'text';
});

const autocomplete = computed(() => {
    if (props.meta === 'login') return 'username';
    if (props.meta === 'password') return 'current-password';
    return 'off';
});

const handleInput = (event) => {
    let value = event.target.value;

    const activeAlphabets = Array.isArray(props.allowedAlphabets)
        ? props.allowedAlphabets
        : props.allowedAlphabets.split(',').map(s => s.trim());

    if (activeAlphabets.length > 0) {
        const patterns = {
            ru: 'а-яА-ЯёЁ',
            en: 'a-zA-Z',
            numeric: '0-9',
            symbols: '!@#$%^&*()_+\\-=\\[\\]{};\':",./<>?\\\\|'
        };

        const activeChars = activeAlphabets
            .map(key => patterns[key])
            .filter(Boolean)
            .join('');

        if (activeChars) {
            const regex = new RegExp(`[^${activeChars}]`, 'g');
            value = value.replace(regex, '');
        }
    }

    event.target.value = value;
    emit('update:modelValue', value);
};
</script>


<template>
    <div class="input-wrapper">
        <input
            :value="Value"
            :type="inputType"
            :placeholder="placeholder"
            :autocomplete="autocomplete"
            @input="handleInput"
            class="custom-input"
        />
    </div>
</template>

<style scoped>
.input-wrapper {
    display: flex;
    flex-direction: column;
}
.custom-input {
    padding: 8px 12px;
    border: 1px solid var(--gray);
    border-radius: 4px;
    outline: none;
}
.custom-input:focus {
    border-color: var(--primary);
}
</style>