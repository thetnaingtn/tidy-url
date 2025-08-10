import { createChannel, createClientFactory, FetchTransport } from "nice-grpc-web";
import { TidyUrlServiceDefinition } from "./types/proto/api/v1/shortener";


const channel = createChannel(
    window.location.origin,
    FetchTransport({
        credentials:"include"
    })
)

const clientFactory = createClientFactory()

export const tidyUrlServiceClient = clientFactory.create(TidyUrlServiceDefinition, channel);