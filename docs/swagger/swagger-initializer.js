window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: "./swagger.yml",
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout",
    onComplete: function () {
      console.info("onComplete completed");
      // const $select = $(".servers select");
      // console.info("$select.length", $select.length);
      $(function () {
        const currentURL = window.location.href;
        console.info("currentURL", currentURL);

        function getHostname(href) {
          const l = document.createElement("a");
          l.href = href;
          return l.hostname;
        }

        const hostname = getHostname(currentURL);
        console.info("hostname", hostname);
        let found = false;
        const $select = $(".servers select");
        console.info("$select.length", $select.length);
        $("body").on("load", ".servers select", function () {
          console.info("select changed");
        });
        const $opts = $(".servers option");
        console.info("$opts len:", $opts.length);
        $opts.each(function (i, o) {
          const $o = $(o);
          const cur = getHostname($o.attr("value"));
          console.info("cur", cur);
          if (cur === hostname) {
            found = true;
          }
        });
        /*if (found) {
            $(".servers option").each(function (i, o) {
                const $o = $(o);
                const cur = getHostname($o.attr("value"));
                if (cur !== hostname) {
                    console.info("removing", cur, "because it's not equal to", hostname);
                    $o.remove();
                }
            });
        }*/
      });
    }
  });

  //</editor-fold>
};
