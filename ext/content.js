chrome.runtime.onMessage.addListener(msg => {
  console.log("Content script received message:", msg);
  if (msg.type === "HOTKEY_EVENT") {
    console.log("YTS", "HOTKEY_EVENT", msg.action)
    if (msg.action === "next_short") {
      const element = document.querySelector('#navigation-button-down button');
      if (element) {
        element.click();
      } else {
        console.warn("Element not found for selector");
      }
    }
  } else if (msg.type === "HOST_ACTION") {
    console.log("Action from host, triggered by the extension:", msg.action);
  }
});

// Start Host Application
function init() {
    if (window.location.href.includes("/shorts/")) {
        console.log("YTShorter: Detected Shorts page, initializing...");
        // Your initialization logic here
        void chrome.runtime.sendMessage({type: "STARTUP_HOST", action: "start_host"}, (response) => {
            if (chrome.runtime.lastError) {
                console.warn("Startup message ignored:", chrome.runtime.lastError.message);
            } else {
                console.log("Host status:", response);
            }
        });
    }
}

// Listen for YouTube's internal navigation event
window.addEventListener("yt-navigate-finish", init);

// Initial run for direct loads
init();
