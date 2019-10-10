/*
 * Copyright 2015-2019 Jeeva Kandasamy (jkandasa@gmail.com)
 * and other contributors as indicated by the @author tags.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.mycontroller.standalone.backup.mixins;

import org.mycontroller.standalone.db.tables.FirmwareType;
import org.mycontroller.standalone.db.tables.FirmwareVersion;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

/**
 * @author Jeeva Kandasamy (jkandasa)
 * @since 1.5.0
 */

@JsonIgnoreProperties({ "fileString", "fileBytes", "fileType", "blockSize", "firmwareName" })
public abstract class MixinFirmware {

    @JsonSerialize(using = SerializerSimpleFirmwareType.class)
    public abstract FirmwareType getType();

    @JsonSerialize(using = SerializerSimpleFirmwareVersion.class)
    public abstract FirmwareVersion getVersion();
}
