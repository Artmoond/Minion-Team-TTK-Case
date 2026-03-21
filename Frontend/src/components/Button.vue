<script setup>
defineProps({
    text: { type: String, default: 'Просто кнопка' },
    icon: { type: String, default: '' },
    disabled: { type: Boolean, default: false }
});

defineEmits(['click']);
</script>


<template>
    <button
        class="custom-button"
        :disabled="disabled"
        @click="$emit('click')"
    >
        <span v-if="icon" class="button-icon">{{ icon }}</span>

        <span class="button-text">
            <slot>{{ text }}</slot>
        </span>
    </button>
</template>


<style scoped>
.custom-button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 10px 20px;
    border: 1px solid var(--gray);
    border-radius: 8px;
    background-color: var(--light);
    cursor: pointer;
    transition: all 0.3s ease;
    position: relative;
}

.custom-button:hover {
    background-color: var(--gray);
}

.button-icon {
    font-size: 16px;
}
</style>