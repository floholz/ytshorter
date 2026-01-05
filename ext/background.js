let port = null;

function ensureNative() {
  if (!port) {
    console.log("Connecting to native host com.floholz.ytshorter");
    port = chrome.runtime.connectNative("com.floholz.ytshorter");
    
    port.onMessage.addListener(msg => {
      console.log("Received from native:", msg);
      chrome.tabs.query({ active: true, url: "https://www.youtube.com/shorts/*" },tabs => {
        for (const tab of tabs) {
            void chrome.tabs.sendMessage(tab.id, msg);
          }
      });
    });

    port.onDisconnect.addListener(() => {
      console.log("Disconnected from native host:", chrome.runtime.lastError);
      port = null;
    });
  }
}

chrome.runtime.onMessage.addListener((msg, sender, sendResponse) => {
  console.log("Received from content script:", msg);
  
  if (msg.type === "STARTUP_HOST") {
    ensureNative();
    sendResponse({ status: "ok", detail: "Native host check initiated" });
  } else if (msg.type === "EXTENSION_ACTION") {
    ensureNative();
    if (port) {
      port.postMessage(msg);
      sendResponse({ status: "sent" });
    } else {
      sendResponse({ status: "error", detail: "Native port not available" });
    }
  }
  return true; // Keeps the message channel open for async responses
});
