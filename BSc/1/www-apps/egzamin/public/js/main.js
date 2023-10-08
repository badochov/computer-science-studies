const main = async () => {
  try {
    const raw = await fetch("/posts");
    const posts = await raw.json();
    const ulEl = document.getElementsByTagName("ul")[0];
    posts.forEach((post) => {
      const li = document.createElement("li");
      li.textContent = post.login_osoby + " : " + post.tresc;
      ulEl.appendChild(li);
    });
  } catch (e) {
    console.error(e);
  }
};

main();
