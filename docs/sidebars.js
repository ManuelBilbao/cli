/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  tutorialSidebar: [
    { type: "autogenerated", dirName: "." },

    {
      type: "category",
      label: "Resources",
      collapsed: true,
      items: [
        {
          type: "link",
          label: "Ignite CLI on Github",
          href: "https://github.com/manuelbilbao/cli",
        },
        {
          type: "link",
          label: "Cosmos SDK Docs",
          href: "https://docs.cosmos.network/",
        },
      ],
    },
  ],

  // But you can create a sidebar manually
  /*
  tutorialSidebar: [
    {
      type: 'category',
      label: 'Tutorial',
      items: ['hello'],
    },
  ],
   */
};

module.exports = sidebars;
