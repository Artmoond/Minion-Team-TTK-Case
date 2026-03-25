<script>
import Button from "@/components/Button.vue";

export default {
    name: 'PageAdmin',
    components: {
        Button
    },
    data() {
        return {
            users: [
                {
                    login: 'user1',
                    fio: 'Иванов Иван Иванович',
                    role: 'Игрок',
                    registrationDate: '2023-01-15'
                },
                {
                    login: 'admin1',
                    fio: 'Петров Петр Петрович',
                    role: 'Администратор',
                    registrationDate: '2023-02-20'
                },
                {
                    login: 'host1',
                    fio: 'Сидоров Сидор Сидорович',
                    role: 'Ведущий',
                    registrationDate: '2023-03-10'
                }
            ]
        }
    }
}
</script>

<template>
    <div class="container mx-auto p-4">
        <h1 class="text-2xl font-bold mb-4">Управление пользователями</h1>
        <table class="min-w-full bg-white border border-gray-300">
            <thead>
                <tr class="bg-gray-100">
                    <th class="py-2 px-4 border-b">Логин</th>
                    <th class="py-2 px-4 border-b">ФИО</th>
                    <th class="py-2 px-4 border-b">Роль</th>
                    <th class="py-2 px-4 border-b">Дата регистрации</th>
                    <th class="py-2 px-4 border-b">Действия</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users" :key="user.login" class="hover:bg-gray-50">
                    <td class="py-2 px-4 border-b">{{ user.login }}</td>
                    <td class="py-2 px-4 border-b">{{ user.fio }}</td>
                    <td class="py-2 px-4 border-b">{{ user.role }}</td>
                    <td class="py-2 px-4 border-b">{{ user.registrationDate }}</td>
                    <td class="py-2 px-4 border-b">
                        <Button 
                            text="Редактировать"
                        />
                        <Button 
                            text="Изменить пароль"
                        />
                        <Button 
                            text="Назначить роли"
                        />
                        <Button 
                            text="Удалить"
                        />
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>