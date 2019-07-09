 code generator tags 
 
// +genclient - generate default client verb functions (create, update, delete, get, list, update, patch, watch and depending on the existence of .Status field in the type the client is generated for also updateStatus).

// +genclient:nonNamespaced - all verb functions are generated without namespace.

// +genclient:onlyVerbs=create,get - only listed verb functions will be generated.

// +genclient:skipVerbs=watch - all default client verb functions will be generated except watch verb.

// +genclient:noStatus - skip generation of updateStatus verb even thought the .Status field exists.

// +genclient:method=Scale,verb=update,subresource=scale,input=k8s.io/api/extensions/v1beta1.Scale,result=k8s.io/api/extensions/v1beta1.Scale - in this case a new function Scale(string, *v1beta.Scale) *v1beta.Scalewill be added to the default client and the body of the function will be based on the update verb. The optional subresource argument will make the generated client function use subresource scale. Using the optional input and result arguments you can override the default type with a custom type. If the import path is not given, the generator will assume the type exists in the same package.

// +groupName=policy.authorization.k8s.io – used in the fake client as the full group name (defaults to the package name).

// +groupGoName=AuthorizationPolicy – a CamelCase Golang identifier to de-conflict groups with non-unique prefixes like policy.authorization.k8s.io and policy.k8s.io. These would lead to two Policy() methods in the clientset otherwise (defaults to the upper-case first segement of the group name).

// +k8s:deepcopy-gen:interfaces tag can and should also be used in cases where you define API types that have fields of some interface type, for example, field SomeInterface. Then // +k8s:deepcopy-gen:interfaces=example.com/pkg/apis/example.SomeInterface will lead to the generation of a DeepCopySomeInterface() SomeInterface method. This allows it to deepcopy those fields in a type-correct way.

// +groupName=example.com defines the fully qualified API group name. If you get that wrong, client-gen will produce wrong code. Be warned that this tag must be in the comment block just above package