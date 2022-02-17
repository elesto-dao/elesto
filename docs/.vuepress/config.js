module.exports = {
    theme: "cosmos",
    title: "Elesto",
    markdown: {
        config: md => {
            md.use(require('markdown-it-plantuml'))
        }
    },
    head: [
      [
        "script",
        /*
        ** Google Tag Manager
        */
        // {
        //   async: true,
        //   src: "https://www.googletagmanager.com/gtag/js?id=G-XL9GNV1KHW",
        // },
      ],
      [
        "script",
        {},
        [
        //   "window.dataLayer = window.dataLayer || [];\nfunction gtag(){dataLayer.push(arguments);}\ngtag('js', new Date());\ngtag('config', 'G-XL9GNV1KHW');",
        ],
      ],
    ],
    themeConfig: {
      logo: {
        src: "/logo.png",
      },
      sidebar: {
        auto: true,
        nav: [
          {
            title: "Resources",
            children: [
              {
                title: "Elesto on Github",
                path: "https://github.com/elesto-dao/elesto",
              },
              {
                title: "Cosmos SDK Docs",
                path: "https://docs.cosmos.network",
              },
            ],
          },
        ],
      },
      topbar: {
        banner: false,
      },
      custom: true,
      footer: {
        question: {
          text:
            "Chat with Elesto and Cosmos SDK developers in <a href='https://github.com/elesto-dao/elesto' target='_blank'>Discord</a>.",
        },
        logo: "/logo.png",
        textLink: {
          text: "Elesto",
          url: "https://github.com/elesto-dao/elesto",
        },
        services: [
          {
            service: "medium",
            url: "https://github.com/elesto-dao/elesto",
          },
          {
            service: "twitter",
            url: "https://github.com/elesto-dao/elesto",
          },
          {
            service: "linkedin",
            url: "https://github.com/elesto-dao/elesto",
          },
          {
            service: "discord",
            url: "https://github.com/elesto-dao/elesto",
          },
          {
            service: "youtube",
            url: "https://github.com/elesto-dao/elesto",
          },
        ],

        smallprint:
          "Random Text",
        links: [
          {
            title: "Documentation",
            children: [
              {
                title: "Cosmos SDK",
                url: "https://docs.cosmos.network",
              },
              {
                title: "Cosmos Hub",
                url: "https://hub.cosmos.network",
              },
              {
                title: "Tendermint Core",
                url: "https://docs.tendermint.com",
              },
            ],
          },
          {
            title: "Community",
            children: [
            //   {
            //     title: "Cosmos blog",
            //     url: "https://blog.cosmos.network",
            //   },
            //   {
            //     title: "Forum",
            //     url: "https://forum.cosmos.network",
            //   },
            //   {
            //     title: "Chat",
            //     url: "https://discord.gg/starport",
            //   },
            ],
          },
          {
            title: "Contributing",
            children: [
              {
                title: "Contributing to the docs",
                url:
                  "https://github.com/cosmos/cosmos-sdk/blob/main/docs/DOCS_README.md",
              },
              {
                title: "Source code on GitHub",
                url: "https://github.com/cosmos/cosmos-sdk/",
              },
            ],
          },
        ],
      },
    },
  };
