import React from "react";
import clsx from "clsx";
import styles from "./styles.module.css";

import Translate, {translate} from '@docusaurus/Translate';

const FeatureList = [
  {
    title: translate({
      id: "homepage.mainBlock1.label",
      message: "Lightweight and Flexible",
      description: 'The homepage icon alt message',
    }),
    Svg: require("@site/static/img/lightweight-and-flexible.svg").default,
    description: translate({
      id: "homepage.mainBlock1.description",
      message: "A Lightweight and Flexible Editor of the Multidimensional Open Gnoseological Active Network (MOGAN) with love by anondigriz and friends in Go.",
      description: 'The homepage icon alt message',
    }),
  },
  {
    title: "Based on the Mivar-based approach",
    Svg: require("@site/static/img/ai.svg").default,
    description: (
      <>
        The Mivar-based approach is a mathematical tool for designing artificial
        intelligence (AI) systems. Mivar (Multidimensional Informational
        Variable Adaptive Reality) was developed by combining production and
        Petri nets.
      </>
    ),
  },
  {
    title: "Tool for creating AI systems",
    Svg: require("@site/static/img/undraw_docusaurus_tree.svg").default,
    description: (
      <>
        The knowledge bases based on the Mivar approach are used for semantic analysis
        and adequate representation of humanitarian epistemological and
        axiological principles in the process of developing AI.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
