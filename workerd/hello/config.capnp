using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
	services = [
		(name = "main", worker = .mainWorker)
	],

	sockets = [
		( name = "http",
			address = "*:8088",
			http = (),
			service = "main"
		),
	]
);

const mainWorker :Workerd.Worker = (
	compatibilityDate = "2022-09-17",

	modules = [
		( name = "worker.mjs", esModule = embed "worker.mjs" ),
	],
);