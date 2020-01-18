                
<p class="summary"><span class="seoSummary">HTTP response status codes indicate whether a specific <a href="/en-US/docs/Web/HTTP">HTTP</a> request has been successfully completed. Responses are grouped in five classes: informational responses, successful responses, redirects, client errors, and servers errors.</span>&nbsp;The below status codes are defined by <a class="external external-icon" href="https://tools.ietf.org/html/rfc2616#section-10" rel="noopener">section 10 of RFC 2616</a>. You can find an updated specification in <a class="external external-icon" href="https://tools.ietf.org/html/rfc7231#section-6.5.1" rel="noopener">RFC 7231</a>.&nbsp;</p>

<h2 id="Information_responses">Information responses<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Information_responses" data-heading-id="Information_responses"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<dl>
 <dt><a href="/en-US/docs/Web/HTTP/Status/100" title="The HTTP 100 Continue informational status response code indicates that everything so far is OK and that the client should continue with the request or ignore it if it is already finished."><code>100 Continue</code></a></dt>
 <dd>This interim response indicates that everything so far is OK and that the client should continue with the request or ignore it if it is already finished.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/101" title="The HTTP&nbsp;101 Switching Protocols response code indicates the protocol the server is switching to as requested by a client which sent the message including the Upgrade request header."><code>101 Switching Protocol</code></a></dt>
 <dd>This code is sent in response to an <a class="new" href="/en-US/docs/Web/HTTP/Headers/Upgrade" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>Upgrade</code></a> request header by the client, and indicates the protocol the server is switching to.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/102" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>102 Processing</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>This code indicates that the server has received and is processing the request, but no response is available yet.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/103" title="The HTTP 103 Early Hints information response status code is primarily intended to be used with the Link header to allow the user agent to start preloading resources while the server is still preparing a response."><code>103 Early Hints</code></a></dt>
 <dd>This status code is primarily intended to be used with the <a href="/en-US/docs/Web/HTTP/Headers/Link" title="The HTTP Link entity-header field provides a means for serialising one or more links in HTTP headers. It is semantically equivalent to the HTML <link> element."><code>Link</code></a> header to allow the user agent to start preloading resources while the server is still preparing a response.</dd>
</dl>

<h2 id="Successful_responses">Successful responses<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Successful_responses" data-heading-id="Successful_responses"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<dl>
 <dt><a href="/en-US/docs/Web/HTTP/Status/200" title="The HTTP 200 OK success status response code indicates that the request has succeeded. A 200 response is cacheable by default."><code>200 OK</code></a></dt>
 <dd>The request has succeeded. The meaning of a success varies depending on the HTTP method:<br>
 GET: The resource has been fetched and is transmitted in the message body.<br>
 HEAD: The entity headers are in the message body.<br>
 PUT or POST: The resource describing the result of the action is transmitted in the message body.<br>
 TRACE: The message body contains the request message as received by the server</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/201" title="The HTTP 201 Created success status response code indicates that the request has succeeded and has led to the creation of a resource. The new resource is effectively created before this response is sent back&nbsp;and the new resource is returned in the body of the message, its location being either the URL of the request, or the content of the Location header."><code>201 Created</code></a></dt>
 <dd>The request has succeeded and a new resource has been created as a result of it. This is typically the response sent after a POST request, or after some PUT requests.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/202" title="The HyperText Transfer Protocol (HTTP) 202 Accepted response status code indicates that the request has been received but not yet acted upon. It is non-committal, meaning that there is no way for the HTTP to later send an asynchronous response indicating the outcome of processing the request. It is intended for cases where another process or server handles the request, or for batch processing."><code>202 Accepted</code></a></dt>
 <dd>The request has been received but not yet acted upon. It is non-committal, meaning that there is no way in HTTP to later send an asynchronous response indicating the outcome of processing the request. It is intended for cases where another process or server handles the request, or for batch processing.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/203" title="The HTTP 203 Non-Authoritative Information response status indicates that the request was successful but the enclosed payload has been modified&nbsp;by a transforming proxy from that of the origin server's 200 (OK) response ."><code>203 Non-Authoritative Information</code></a></dt>
 <dd>This response code means returned meta-information set is not exact set as available from the origin server, but collected from a local or a third party copy. Except this condition, 200 OK response should be preferred instead of this response.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/204" title="The HTTP 204 No Content success status response code indicates that the request has succeeded, but that the client doesn't need to go away from its current page. A 204 response is cacheable by default. An ETag header is included in such a response."><code>204 No Content</code></a></dt>
 <dd>There is no content to send for this request, but the headers may be useful. The user-agent may update its cached headers for this resource with the new ones.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/205" title="The HTTP 205 Reset Content response status tells the client to reset the document view, so for example to clear the content of a form, reset a canvas state, or to refresh the UI."><code>205 Reset Content</code></a></dt>
 <dd>This response code is sent after accomplishing request to tell user agent reset document view which sent this request.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/206" title="The HTTP 206 Partial Content success status response code indicates that the request has succeeded and has the body contains the requested ranges of data, as described in the Range header of the request."><code>206 Partial Content</code></a></dt>
 <dd>This response code is used because of range header sent by the client to separate download into multiple streams.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/207" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>207 Multi-Status</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>A Multi-Status response conveys information about multiple resources in situations where multiple status codes might be appropriate.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/208" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>208 Multi-Status</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>Used inside a DAV: propstat response element to avoid enumerating the internal members of multiple bindings to the same collection repeatedly.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/226" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>226 IM Used</code></a> (<a class="external external-icon" href="https://tools.ietf.org/html/rfc3229" rel="noopener">HTTP Delta encoding</a>)</dt>
 <dd>The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.</dd>
</dl>

<h2 id="Redirection_messages">Redirection messages<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Redirection_messages" data-heading-id="Redirection_messages"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<dl>
 <dt><a href="/en-US/docs/Web/HTTP/Status/300" title="The HTTP 300 Multiple Choices redirect status response code indicates that the request has more than one possible responses. The user-agent or the user should choose one of them. As there is no standardized way of choosing one of the responses, this response code is very rarely used."><code>300 Multiple Choice</code></a></dt>
 <dd>The request has more than one possible response. The user-agent or user should choose one of them. There is no standardized way of choosing&nbsp;one of the responses.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/301" title="The HyperText Transfer Protocol (HTTP) 301 Moved Permanently redirect status response code indicates that the resource requested has been definitively moved to the URL given by the Location headers. A browser redirects to this page and search engines update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is sent to the new URL)."><code>301 Moved Permanently</code></a></dt>
 <dd>This response code means that the URI of the requested resource has been changed permanently. Probably, the new URI would be given in the response.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/302" title="The HyperText Transfer Protocol (HTTP) 302 Found redirect status response code indicates that the resource requested has been temporarily moved to the URL given by the Location header. A browser redirects to this page but search engines don't update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is not sent to the new URL)."><code>302 Found</code></a></dt>
 <dd>This response code means that the URI of requested resource has been changed <em>temporarily</em>. New changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/303" title="The HyperText Transfer Protocol (HTTP) 303 See Other redirect status response code indicates that the redirects don't link to the newly uploaded resources but to another page, like a confirmation page or an upload progress page. This response code is usually sent back as a result of PUT or POST. The method used to display this redirected page is always GET."><code>303 See Other</code></a></dt>
 <dd>The server sent this response to direct the&nbsp;client to get the requested resource at another URI with a&nbsp;GET request.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/304" title="The HTTP 304 Not Modified client redirection response code indicates that there is no need to retransmit the requested resources. It is an implicit redirection to a cached resource. This happens when the request method is safe, like a GET or a HEAD request, or when the request is conditional and uses a If-None-Match or a If-Modified-Since header."><code>304 Not Modified</code></a></dt>
 <dd>This is used for caching purposes. It tells the client that the response has not been modified, so the client can continue to use the same cached version of the response.</dd>
 <dt><code>305 Use Proxy</code> <span title="This deprecated API should no longer be used, but will probably still work."><i class="icon-thumbs-down-alt"> </i></span></dt>
 <dd>Was defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.</dd>
 <dt><code>306 unused</code></dt>
 <dd>This response code is no longer used, it is just reserved currently. It was used in a previous version of the HTTP 1.1 specification.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/307" title="The method and the body of the original request are reused to perform the redirected request. In the cases where you want the method used to be changed to GET, use 303 See Other instead. This is useful when you want to give an answer to a PUT method that is not the uploaded resources, but a confirmation message (like &quot;You successfully uploaded XYZ&quot;)."><code>307 Temporary Redirect</code></a></dt>
 <dd>The server sends this response to direct the&nbsp;client to get the requested resource at another URI with same method that was used in the prior request. This has the same semantics as the <code>302 Found</code> HTTP response code, with the exception that the user agent <em>must not</em> change the HTTP method used: If a <code>POST</code> was used in the first request, a <code>POST</code> must be used in the second request.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/308" title="The request method and the body will not be altered, whereas 301 may incorrectly sometimes be changed to a GET method."><code>308 Permanent Redirect</code></a></dt>
 <dd>This means that the resource is now permanently located at another URI, specified by the <code>Location:</code> HTTP Response header. This has the same semantics as the <code>301 Moved Permanently</code> HTTP response code, with the exception that the user agent <em>must not</em> change the HTTP method used: If a <code>POST</code> was used in the first request, a <code>POST</code> must be used in the second request.</dd>
</dl>

<h2 id="Client_error_responses">Client error responses<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Client_error_responses" data-heading-id="Client_error_responses"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<dl>
 <dt><a href="/en-US/docs/Web/HTTP/Status/400" title="The HyperText Transfer Protocol (HTTP) 400 Bad Request response status code indicates that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)."><code>400 Bad Request</code></a></dt>
 <dd>This response means that server could not understand the request due to invalid syntax.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/401" title="The HTTP 401 Unauthorized client error status response code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource."><code>401 Unauthorized</code></a></dt>
 <dd>Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.</dd>
 <dt><code>402 Payment Required</code></dt>
 <dd>This response code is reserved for future use. Initial aim for creating this code was using it for digital payment systems however this is not used currently.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/403" title="The HTTP 403 Forbidden client error status response code indicates that the server understood the request but refuses to authorize it."><code>403 Forbidden</code></a></dt>
 <dd>The client&nbsp;does not have access rights to the content, i.e. they are unauthorized,&nbsp;so server is rejecting to give proper response. Unlike 401, the client's identity is known to the server.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/404" title="The HTTP 404 Not Found client error response code indicates that the&nbsp;server can't&nbsp;find the requested resource. Links which lead to a 404 page are often called broken or dead links,&nbsp;and can be&nbsp;subject&nbsp;to link rot."><code>404 Not Found</code></a></dt>
 <dd>The server can not find requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client. This response code is probably the&nbsp;most famous one due to its frequent occurence on the&nbsp;web.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/405" title="The HyperText Transfer Protocol (HTTP) 405 Method Not Allowed response status code indicates that the request method is known by the server but is not supported by the target resource."><code>405 Method Not Allowed</code></a></dt>
 <dd>The request method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods,&nbsp;<code>GET</code> and <code>HEAD</code>, must never be disabled and should not return this error code.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/406" title="The HyperText Transfer Protocol (HTTP) 406 Not Acceptable client error response code indicates that the server cannot produce a response&nbsp;matching the list of acceptable values defined in the request's&nbsp;proactive content negotiation headers,&nbsp;and that the server is unwilling to supply a default representation."><code>406 Not Acceptable</code></a></dt>
 <dd>This response is sent when the web server, after performing <a href="/en-US/docs/HTTP/Content_negotiation#Server-driven_negotiation">server-driven content negotiation</a>, doesn't find any content following the criteria given by the user agent.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/407" title="The HTTP 407 Proxy Authentication Required  client error status response code indicates that the request has not been applied because it lacks valid authentication credentials for a proxy server that is between the browser and the server that can access the requested resource."><code>407 Proxy Authentication Required</code></a></dt>
 <dd>This is similar to 401 but authentication is needed to be done by a proxy.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/408" title="The HyperText Transfer Protocol (HTTP) 408 Request Timeout response status code&nbsp;means that the server would like to shut down this unused connection. It is sent on an idle connection by some servers, even without any previous request by the client."><code>408 Request Timeout</code></a></dt>
 <dd>This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/409" title="The HTTP 409 Conflict response status code indicates a request conflict with current state of the server."><code>409 Conflict</code></a></dt>
 <dd>This response is&nbsp;sent when a request conflicts with the current state of the server.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/410" title="The HyperText Transfer Protocol (HTTP) 410 Gone client error response code indicates that access to the target resource is no longer available at the origin server and that this condition is likely to be permanent."><code>410 Gone</code></a></dt>
 <dd>This response would be sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for "limited-time, promotional services". APIs should not feel compelled to indicate resources that have been deleted with this status code.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/411" title="The HyperText Transfer Protocol (HTTP) 411 Length Required client error response code indicates that the server refuses to accept the request without a defined Content-Length header."><code>411 Length Required</code></a></dt>
 <dd>Server rejected the request because the <code>Content-Length</code> header field is not defined and the server requires it.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/412" title="The HyperText Transfer Protocol (HTTP) 412 Precondition Failed client error response code indicates that access to the target resource has been denied. This happens with conditional requests on methods other than GET or HEAD when the condition defined by the If-Unmodified-Since or If-None-Match headers is not fulfilled. In that case, the request, usually an upload or a modification of a resource, cannot be made and this error response is sent back."><code>412 Precondition Failed</code></a></dt>
 <dd>The client has indicated preconditions in its headers which the server does not meet.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/413" title="The HTTP 413 Payload Too Large response status code indicates that the request entity is larger than limits defined by server; the server might close the connection or return a Retry-After header field."><code>413 Payload Too Large</code></a></dt>
 <dd>Request entity is larger than limits defined by server; the server might close the connection or return an <code>Retry-After</code> header field.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/414" title="The HTTP 414 URI Too Long response status code indicates that the URI requested by the client is longer than the server is willing to interpret."><code>414 URI Too Long</code></a></dt>
 <dd>The URI requested by the client is longer than the server is willing to interpret.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/415" title="The HTTP 415 Unsupported Media Type client error response code indicates that the server refuses to accept the request because the payload format is in an unsupported format."><code>415 Unsupported Media Type</code></a></dt>
 <dd>The media format of the requested data is not supported by the server, so the server is rejecting the request.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/416" title="The HyperText Transfer Protocol (HTTP) 416 Range Not Satisfiable error response code indicates that a server cannot serve the requested ranges. The most likely reason is that the document doesn't contain such ranges, or that the Range header value, though syntactically correct, doesn't make sense."><code>416 Requested Range Not Satisfiable</code></a></dt>
 <dd>The range specified by the <code>Range</code> header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/417" title="The HTTP 417 Expectation Failed client error response code indicates that the expectation given in the request's Expect header could not be met."><code>417 Expectation Failed</code></a></dt>
 <dd>This response code means the expectation indicated by the <code>Expect</code> request header field can't be met by the server.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/418" title="The HTTP 418 I'm a teapot client error response code indicates that the server refuses to brew coffee because it is a teapot. This error is a reference of Hyper Text Coffee Pot Control Protocol which was an April Fools' joke in 1998."><code>418 I'm a teapot</code></a></dt>
 <dd>The server refuses the attempt to brew coffee with a teapot.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/421" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>421 Misdirected Request</code></a></dt>
 <dd>The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/422" title="The HyperText Transfer Protocol (HTTP) 422 Unprocessable Entity response status code indicates that the server understands the content type of the request entity, and the syntax of the request entity is correct, but it was unable to process the contained instructions."><code>422 Unprocessable Entity</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>The request was well-formed but was unable to be followed due to semantic errors.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/423" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>423 Locked</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>The resource that is being accessed is locked.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/424" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>424 Failed Dependency</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>The request failed due to failure of a previous request.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/425" title="The HyperText Transfer Protocol (HTTP) 425 Too Early response status code indicates that the server is unwilling to risk processing a request that might be replayed, which creates the potential for a replay attack."><code>425 Too Early</code></a></dt>
 <dd>Indicates that the server is unwilling to risk processing a request that might be replayed.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/426" title="The HTTP 426 Upgrade Required client error response code indicates that the server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol."><code>426 Upgrade Required</code></a></dt>
 <dd>The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server sends an <a class="new" href="/en-US/docs/Web/HTTP/Headers/Upgrade" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>Upgrade</code></a> header in a 426 response to indicate the required protocol(s).</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/428" title="The HTTP 428 Precondition Required response status code indicates that the server requires the request to be conditional."><code>428 Precondition Required</code></a></dt>
 <dd>The origin server requires the request to be conditional. Intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/429" title="The HTTP 429 Too Many Requests response status code indicates the user has sent too many requests in a given amount of time (&quot;rate limiting&quot;)."><code>429 Too Many Requests</code></a></dt>
 <dd>The user has sent too many requests in a given amount of time ("rate limiting").</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/431" title="The HTTP 431 Request Header Fields Too Large response status code indicates that the server is unwilling to process the request because its header fields are too large. The request may be resubmitted after reducing the size of the request header fields."><code>431 Request Header Fields Too Large</code></a></dt>
 <dd>The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/451" title="The HyperText Transfer Protocol (HTTP) 451 Unavailable For Legal Reasons client error response code indicates that the user requested a resource that is not available due to legal reasons, such as a web page for which a legal action has been issued."><code>451 Unavailable For Legal Reasons</code></a></dt>
 <dd>The user requests an illegal resource, such as a web page censored by a government.</dd>
</dl>

<h2 id="Server_error_responses">Server error responses<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Server_error_responses" data-heading-id="Server_error_responses"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<dl>
 <dt><a href="/en-US/docs/Web/HTTP/Status/500" title="The HyperText Transfer Protocol (HTTP) 500 Internal Server Error server error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request."><code>500 Internal Server Error</code></a></dt>
 <dd>The server has encountered a situation it doesn't know how to handle.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/501" title="The HyperText Transfer Protocol (HTTP) 501 Not Implemented server error response code indicates that the server does not support the functionality required to fulfill the request. This is the appropriate response when the server does not recognize the request method and is not capable of supporting it for any resource. The only request methods that servers are required to support (and therefore that must not return this code) are GET and HEAD."><code>501 Not Implemented</code></a></dt>
 <dd>The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are <code>GET</code>&nbsp;and <code>HEAD</code>.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/502" title="The HyperText Transfer Protocol (HTTP) 502 Bad Gateway server error response code indicates that the server, while acting as a gateway or proxy, received an invalid response from the upstream server."><code>502 Bad Gateway</code></a></dt>
 <dd>This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/503" title="The HyperText Transfer Protocol (HTTP) 503 Service Unavailable server error response code indicates that the server is not ready to handle the request."><code>503 Service Unavailable</code></a></dt>
 <dd>The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This responses should be used for temporary conditions and the <code>Retry-After:</code> HTTP&nbsp;header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/504" title="The HyperText Transfer Protocol (HTTP) 504 Gateway Timeout server error response code indicates that the server, while acting as a gateway or proxy, did not get a response in time from the upstream&nbsp;server that it needed in order to complete the request."><code>504 Gateway Timeout</code></a></dt>
 <dd>This error response is given when the server is acting as a gateway and cannot get a response in time.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/505" title="The HyperText Transfer Protocol (HTTP) 505 HTTP Version Not Supported response status code indicates that the HTTP version used in the request is not supported by the server."><code>505 HTTP Version Not Supported</code></a></dt>
 <dd>The HTTP version used in the request is not supported by the server.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/506" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>506 Variant Also Negotiates</code></a></dt>
 <dd>The server has an internal configuration error: transparent content negotiation for the request results in a circular reference.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/507" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>507 Insufficient Storage</code></a></dt>
 <dd>The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/508" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>508 Loop Detected</code></a> (<a class="glossaryLink" href="/en-US/docs/Glossary/WebDAV" title="WebDAV: WebDAV  (Web Distributed Authoring and Versioning) is an HTTP Extension that lets web developers update their content remotely from a client.">WebDAV</a>)</dt>
 <dd>The server detected an infinite loop while processing the request.</dd>
 <dt><a class="new" href="/en-US/docs/Web/HTTP/Status/510" rel="nofollow" title="The documentation about this has not yet been written; please consider contributing!"><code>510 Not Extended</code></a></dt>
 <dd>Further extensions to the request are required for the server to fulfill it.</dd>
 <dt><a href="/en-US/docs/Web/HTTP/Status/511" title="The HTTP 511 Network Authentication Required response status code indicates that the client needs to authenticate to gain network access."><code>511 Network Authentication Required</code></a></dt>
 <dd>The 511 status code indicates that the client needs to authenticate to gain network access.</dd>
</dl>

<h2 id="Browser_compatibility">Browser compatibility<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#Browser_compatibility" data-heading-id="Browser_compatibility"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<div class="hidden">The compatibility table in this page is generated from structured data. If you'd like to contribute to the data, please check out <a class="external external-icon" href="https://github.com/mdn/browser-compat-data" rel="noopener">https://github.com/mdn/browser-compat-data</a> and send us a pull request.</div>

<p></p><div class="bc-data"><a class="bc-github-link external external-icon" href="https://github.com/mdn/browser-compat-data" rel="noopener">Update compatibility data on GitHub</a><table class="bc-table bc-table-web"><thead><tr class="bc-platforms"><td></td><th class="bc-platform-desktop" colspan="6"><span>Desktop</span></th><th class="bc-platform-mobile" colspan="7"><span>Mobile</span></th></tr><tr class="bc-browsers"><td></td><th class="bc-browser-chrome"><span class="bc-head-txt-label bc-head-icon-chrome">Chrome</span></th><th class="bc-browser-edge"><span class="bc-head-txt-label bc-head-icon-edge">Edge</span></th><th class="bc-browser-firefox"><span class="bc-head-txt-label bc-head-icon-firefox">Firefox</span></th><th class="bc-browser-ie"><span class="bc-head-txt-label bc-head-icon-ie">Internet Explorer</span></th><th class="bc-browser-opera"><span class="bc-head-txt-label bc-head-icon-opera">Opera</span></th><th class="bc-browser-safari"><span class="bc-head-txt-label bc-head-icon-safari">Safari</span></th><th class="bc-browser-webview_android"><span class="bc-head-txt-label bc-head-icon-webview_android">Android webview</span></th><th class="bc-browser-chrome_android"><span class="bc-head-txt-label bc-head-icon-chrome_android">Chrome for Android</span></th><th class="bc-browser-edge_mobile"><span class="bc-head-txt-label bc-head-icon-edge_mobile">Edge Mobile</span></th><th class="bc-browser-firefox_android"><span class="bc-head-txt-label bc-head-icon-firefox_android">Firefox for Android</span></th><th class="bc-browser-opera_android"><span class="bc-head-txt-label bc-head-icon-opera_android">Opera for Android</span></th><th class="bc-browser-safari_ios"><span class="bc-head-txt-label bc-head-icon-safari_ios">Safari on iOS</span></th><th class="bc-browser-samsunginternet_android"><span class="bc-head-txt-label bc-head-icon-samsunginternet_android">Samsung Internet</span></th></tr></thead><tbody><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/100"><code>100</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/200"><code>200</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/201"><code>201</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/204"><code>204</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/206"><code>206</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/301"><code>301</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/302"><code>302</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/303"><code>303</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/304"><code>304</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/307"><code>307</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/308"><code>308</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie bc-has-history" tabindex="0" aria-expanded="false" aria-controls="bc-history-1"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes<div class="bc-icons"><abbr class="only-icon" title="See implementation notes"><span>Notes</span><i class="ic-footnote"></i></abbr> </div><button title="Open implementation notes" class="bc-history-link only-icon" tabindex="-1"><span>Open</span><i class="ic-history" aria-hidden="true"></i></button><section class="bc-history" id="bc-history-1"><dl><dt class="bc-supports-yes bc-supports"><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes<div class="bc-icons"><abbr class="only-icon" title="See implementation notes"><span>Notes</span><i class="ic-footnote"></i></abbr> </div></dt><dd><abbr class="only-icon" title="See implementation notes"><span>Notes</span><i class="ic-footnote"></i></abbr> Does not work on Windows 7 and Windows 8.1.</dd></dl></section></td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/401"><code>401</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/403"><code>403</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/404"><code>404</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/406"><code>406</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/407"><code>407</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/410"><code>410</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/412"><code>412</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/416"><code>416</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/418"><code>418</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/425"><code>425</code></a></th><td class="bc-supports-unknown bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              58</td><td class="bc-supports-unknown bc-browser-ie"><span class="bc-browser-name">IE</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              58</td><td class="bc-supports-unknown bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td><td class="bc-supports-unknown bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr title="Compatibility unknown; please update this.">
                ?
              </abbr></td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/451"><code>451</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/500"><code>500</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/501"><code>501</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/502"><code>502</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/503"><code>503</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr><tr><th scope="row"><a href="/en-US/docs/Web/HTTP/Status/504"><code>504</code></a></th><td class="bc-supports-yes bc-browser-chrome"><span class="bc-browser-name">Chrome</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge"><span class="bc-browser-name">Edge</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox"><span class="bc-browser-name">Firefox</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-ie"><span class="bc-browser-name">IE</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera"><span class="bc-browser-name">Opera</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari"><span class="bc-browser-name">Safari</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-webview_android"><span class="bc-browser-name">WebView Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-chrome_android"><span class="bc-browser-name">Chrome Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-edge_mobile"><span class="bc-browser-name">Edge Mobile</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-firefox_android"><span class="bc-browser-name">Firefox Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-opera_android"><span class="bc-browser-name">Opera Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-safari_ios"><span class="bc-browser-name">Safari iOS</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td><td class="bc-supports-yes bc-browser-samsunginternet_android"><span class="bc-browser-name">Samsung Internet Android</span><abbr class="bc-level-yes only-icon" title="Full support">
                <span>Full support</span>
              </abbr>
              Yes</td></tr></tbody></table><section class="bc-legend" id="sect2"><h3 class="offscreen highlight-spanned" id="Legend"><span class="highlight-span">Legend</span></h3><dl><dt><span class="bc-supports-yes bc-supports">
                <abbr class="bc-level bc-level-yes only-icon" title="Full support">
                 <span>Full support</span>
                 &nbsp;
                </abbr></span></dt><dd>Full support</dd><dt><span class="bc-supports-unknown bc-supports">
                <abbr class="bc-level bc-level-unknown only-icon" title="Compatibility unknown">
                 <span>Compatibility unknown</span>
                 &nbsp;
                </abbr></span></dt><dd>Compatibility unknown</dd><dt><abbr class="only-icon" title="See implementation notes."><span>See implementation notes.</span><i class="ic-footnote"></i></abbr></dt><dd>See implementation notes.</dd></dl></section></div><p></p>

<h2 id="See_also">See also<a class="local-anchor" href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#See_also" data-heading-id="See_also"><svg class="icon icon-link" version="1.1" width="24" height="28" viewBox="0 0 512 512" aria-hidden="true" focusable="false"><path d="M326.612 185.391c59.747 59.809 58.927 155.698.36 214.59-.11.12-.24.25-.36.37l-67.2 67.2c-59.27 59.27-155.699 59.262-214.96 0-59.27-59.26-59.27-155.7 0-214.96l37.106-37.106c9.84-9.84 26.786-3.3 27.294 10.606.648 17.722 3.826 35.527 9.69 52.721 1.986 5.822.567 12.262-3.783 16.612l-13.087 13.087c-28.026 28.026-28.905 73.66-1.155 101.96 28.024 28.579 74.086 28.749 102.325.51l67.2-67.19c28.191-28.191 28.073-73.757 0-101.83-3.701-3.694-7.429-6.564-10.341-8.569a16.037 16.037 0 0 1-6.947-12.606c-.396-10.567 3.348-21.456 11.698-29.806l21.054-21.055c5.521-5.521 14.182-6.199 20.584-1.731a152.482 152.482 0 0 1 20.522 17.197zM467.547 44.449c-59.261-59.262-155.69-59.27-214.96 0l-67.2 67.2c-.12.12-.25.25-.36.37-58.566 58.892-59.387 154.781.36 214.59a152.454 152.454 0 0 0 20.521 17.196c6.402 4.468 15.064 3.789 20.584-1.731l21.054-21.055c8.35-8.35 12.094-19.239 11.698-29.806a16.037 16.037 0 0 0-6.947-12.606c-2.912-2.005-6.64-4.875-10.341-8.569-28.073-28.073-28.191-73.639 0-101.83l67.2-67.19c28.239-28.239 74.3-28.069 102.325.51 27.75 28.3 26.872 73.934-1.155 101.96l-13.087 13.087c-4.35 4.35-5.769 10.79-3.783 16.612 5.864 17.194 9.042 34.999 9.69 52.721.509 13.906 17.454 20.446 27.294 10.606l37.106-37.106c59.271-59.259 59.271-155.699.001-214.959z"></path></svg><span>Section</span></a></h2>

<ul>
 <li><a class="external external-icon" href="https://en.wikipedia.org/wiki/List_of_HTTP_status_codes" rel="noopener">List of HTTP status codes on Wikipedia</a></li>
 <li><a class="external external-icon" href="http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml" rel="noopener">IANA official registry of HTTP status codes</a></li>
</ul>
                  
                
