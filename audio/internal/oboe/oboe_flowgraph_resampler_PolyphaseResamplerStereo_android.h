/*
 * Copyright 2019 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef OBOE_POLYPHASE_RESAMPLER_STEREO_H
#define OBOE_POLYPHASE_RESAMPLER_STEREO_H

#include <sys/types.h>
#include <unistd.h>
#include "oboe_flowgraph_resampler_PolyphaseResampler_android.h"

namespace resampler {

class PolyphaseResamplerStereo : public PolyphaseResampler {
public:
    explicit PolyphaseResamplerStereo(const MultiChannelResampler::Builder &builder);

    virtual ~PolyphaseResamplerStereo() = default;

    void writeFrame(const float *frame) override;

    void readFrame(float *frame) override;
};

}

#endif //OBOE_POLYPHASE_RESAMPLER_STEREO_H