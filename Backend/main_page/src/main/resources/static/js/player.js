// src/main/resources/static/js/player.js

// Ждем полной загрузки страницы
document.addEventListener('DOMContentLoaded', function() {

    // Получаем элементы
    const audioElement = document.getElementById('audio-player');
    const statusElement = document.getElementById('status');

    // Проверяем, существует ли элемент
    if (!audioElement) {
        console.error('Аудио элемент не найден');
        return;
    }

    // Инициализируем Plyr с настройками
    const player = new Plyr(audioElement, {
        // Элементы управления
        controls: [
            'play',           // Воспроизведение/пауза
            'progress',       // Прогресс-бар
            'current-time',   // Текущее время
            'duration',       // Общая длительность
            'mute',           // Отключение звука
            'volume',         // Ползунок громкости
            'settings',       // Настройки (скорость)
            'download'        // Кнопка скачать
        ],

        // Настройки скорости воспроизведения
        speed: {
            selected: 1,
            options: [0.5, 0.75, 1, 1.25, 1.5, 2]
        },

        // Автовоспроизведение (может быть заблокировано)
        autoplay: true,

        // Не сбрасывать при окончании
        resetOnEnd: false,

        // Русская локализация
        i18n: {
            play: 'Воспроизвести',
            pause: 'Пауза',
            mute: 'Отключить звук',
            unmute: 'Включить звук',
            settings: 'Настройки',
            speed: 'Скорость',
            normal: 'Обычная',
            quality: 'Качество',
            loop: 'Повтор'
        },

        // Отладка
        debug: false
    });

    // ===== ОБРАБОТЧИКИ СОБЫТИЙ =====

    // Начало воспроизведения
    player.on('play', () => {
        console.log('▶ Воспроизведение началось');
        statusElement.innerHTML = '🎵 Играет...';
        statusElement.style.color = '#4CAF50';
    });

    // Пауза
    player.on('pause', () => {
        console.log('⏸ Пауза');
        statusElement.innerHTML = '⏸ На паузе';
        statusElement.style.color = '#FF9800';
    });

    // Окончание
    player.on('ended', () => {
        console.log('⏹ Воспроизведение завершено');
        statusElement.innerHTML = '⏹ Остановлено';
        statusElement.style.color = '#666';
    });

    // Ошибка
    player.on('error', (event) => {
        console.error('❌ Ошибка плеера:', event.detail);
        statusElement.innerHTML = '❌ Ошибка потока. Проверьте наличие аудиофайлов';
        statusElement.style.color = '#f44336';
    });

    // Буферизация
    player.on('waiting', () => {
        console.log('⏳ Буферизация...');
        statusElement.innerHTML = '⏳ Буферизация...';
        statusElement.classList.add('loading');
    });

    // Готовность
    player.on('canplay', () => {
        console.log('✅ Поток готов');
        statusElement.innerHTML = '✅ Готов к воспроизведению';
        statusElement.style.color = '#667eea';
        statusElement.classList.remove('loading');
    });

    // Обновление времени (для отладки)
    player.on('timeupdate', (event) => {
        const currentTime = event.detail.plyr.currentTime;
        // Можно использовать для обновления чего-либо
        // console.log('Текущее время:', currentTime);
    });

    // Изменение громкости
    player.on('volumechange', (event) => {
        const volume = event.detail.plyr.volume;
        console.log('Громкость изменена:', Math.round(volume * 100) + '%');
    });

    // Готовность плеера
    player.on('ready', () => {
        console.log('🎬 Плеер Plyr готов к работе');

        // Попытка автовоспроизведения (если заблокировано, показываем подсказку)
        setTimeout(() => {
            if (player.paused) {
                statusElement.innerHTML = '▶ Нажмите play для воспроизведения';
                statusElement.style.color = '#FF9800';
            }
        }, 1000);
    });

    // Обработка потери соединения
    audioElement.addEventListener('stalled', () => {
        console.warn('Поток приостановлен, попытка переподключения...');
        statusElement.innerHTML = '🔄 Переподключение...';
        statusElement.style.color = '#FF9800';

        // Попытка перезагрузки потока
        setTimeout(() => {
            audioElement.load();
        }, 3000);
    });

    // Обработка успешного подключения
    audioElement.addEventListener('canplaythrough', () => {
        console.log('Поток загружен, можно воспроизводить без прерываний');
    });

    console.log('Плеер инициализирован');
});