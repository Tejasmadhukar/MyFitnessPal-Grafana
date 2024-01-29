const theme_switch_icon = document.getElementById("theme_switch_icon");
const github_icon = document.getElementById("github_icon");
const file_name_container = document.getElementById("file_name_container");
const file_choose_button = document.getElementById("file_choose_button");
const file_submit_button = document.getElementById("file_submit_button");
const error_container = document.getElementById("any-errors");

const changeTheme = () => {
  if (document.documentElement.className == "dark") {
    document.documentElement.classList.remove("dark");
    document.documentElement.classList.add("light");
  } else {
    document.documentElement.classList.remove("light");
    document.documentElement.classList.add("dark");
  }
};

const changeIconTheme = () => {
  const xmlns = "http://www.w3.org/2000/svg";

  if (document.documentElement.className == "dark") {
    const path1 = document.createElementNS(xmlns, "path");
    path1.setAttributeNS(
      null,
      "d",
      "M21.53 15.93c-.16-.27-.61-.69-1.73-.49a8.46 8.46 0 01-1.88.13 8.409 8.409 0 01-5.91-2.82 8.068 8.068 0 01-1.44-8.66c.44-1.01.13-1.54-.09-1.76s-.77-.55-1.83-.11a10.318 10.318 0 00-6.32 10.21 10.475 10.475 0 007.04 8.99 10 10 0 002.89.55c.16.01.32.02.48.02a10.5 10.5 0 008.47-4.27c.67-.93.49-1.519.32-1.79z",
    );
    path1.setAttributeNS(null, "fill", "currentColor");
    github_icon.setAttributeNS(null, "fill", "currentColor");

    theme_switch_icon.replaceChildren(path1);
  } else {
    const ele1 = document.createElementNS(xmlns, "g");
    ele1.setAttributeNS(null, "fill", "#f3f4f6");
    github_icon.setAttributeNS(null, "fill", "#f3f4f6");

    const path1 = document.createElementNS(xmlns, "path");
    path1.setAttributeNS(null, "d", "M19 12a7 7 0 11-7-7 7 7 0 017 7z");
    const path2 = document.createElementNS(xmlns, "path");
    path2.setAttributeNS(
      null,
      "d",
      "M12 22.96a.969.969 0 01-1-.96v-.08a1 1 0 012 0 1.038 1.038 0 01-1 1.04zm7.14-2.82a1.024 1.024 0 01-.71-.29l-.13-.13a1 1 0 011.41-1.41l.13.13a1 1 0 010 1.41.984.984 0 01-.7.29zm-14.28 0a1.024 1.024 0 01-.71-.29 1 1 0 010-1.41l.13-.13a1 1 0 011.41 1.41l-.13.13a1 1 0 01-.7.29zM22 13h-.08a1 1 0 010-2 1.038 1.038 0 011.04 1 .969.969 0 01-.96 1zM2.08 13H2a1 1 0 010-2 1.038 1.038 0 011.04 1 .969.969 0 01-.96 1zm16.93-7.01a1.024 1.024 0 01-.71-.29 1 1 0 010-1.41l.13-.13a1 1 0 011.41 1.41l-.13.13a.984.984 0 01-.7.29zm-14.02 0a1.024 1.024 0 01-.71-.29l-.13-.14a1 1 0 011.41-1.41l.13.13a1 1 0 010 1.41.97.97 0 01-.7.3zM12 3.04a.969.969 0 01-1-.96V2a1 1 0 012 0 1.038 1.038 0 01-1 1.04z",
    );

    ele1.appendChild(path1);
    ele1.appendChild(path2);

    theme_switch_icon.replaceChildren(ele1);
  }

  changeTheme();
};

const updateFileName = (event) => {
  file_name_container.children[0].textContent = event.target.files[0].name;
  file_submit_button.style.display = "inline-flex";
  error_container.textContent = "";
};

theme_switch_icon.addEventListener("click", changeIconTheme);
file_choose_button.addEventListener("change", updateFileName);
