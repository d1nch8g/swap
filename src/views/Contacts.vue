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
        Почта: {{ email }}
    </p>
    <p>
        Telegram: {{ telegram }}
    </p>
</template>

<style></style>
