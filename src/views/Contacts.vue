<script>
export default {
    data() {
        return {
            email: "",
            telegram: ""
        }
    },
    async mounted() {
        let response = await fetch(`/api/info`, {
            method: "GET",
            headers: {}
        });

        if (response.ok) {
            let data = await response.json();
            this.email = data.email;
            this.telegram = data.telegram;
        }
    }
}
</script>

<template>
    <title>Контакты</title>
    <meta property="og:title" content="Контакты платформы" />

    <meta name="description"
        content="На данной странице вы можете найти контактный адрес почты и телеграм для обращения по теме созданных заявок." />
    <meta property="og:description"
        content="На данной странице вы можете найти контактный адрес почты и телеграм для обращения по теме созданных заявок." />

    <meta property="og:image" content="../assets/logo.svg" />

    <h1>Контакты</h1>
    <p>
        По любым вопросам связанным с исполнением заявок всегда можно обратиться на почту и в телеграм.
    </p>
    <p>
        Почта: <a :href="'mailto:' + email">{{ email }}</a>
    </p>
    <p>
        Telegram: <a :href="'https://t.me/' + telegram">@{{ telegram }}</a>
    </p>
    <br>
    <p>
        Часы работы: 8:00-24:00, Москва - UTC/GMT +3 hours
    </p>
</template>

<style></style>
